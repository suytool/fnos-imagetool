package main

import (
	"encoding/json"
	"fmt"
	"image"
	"image/color"
	"math"
	"net/http"
	"strconv"
	"strings"

	"github.com/disintegration/imaging"
)

// CellView 单个格子的图片视图调整（缩放/平移）
type CellView struct {
	Zoom float64 `json:"zoom"` // 缩放倍数 0.5-5，1 表示默认 cover 填充
	DX   float64 `json:"dx"`   // 水平偏移百分比（相对可见窗口），-100~100
	DY   float64 `json:"dy"`   // 垂直偏移百分比（相对可见窗口），-100~100
}

// handleCollage /api/collage
// 表单字段：
//   - images[]: 多张图片文件
//   - mode: 模板模式（classic9/2/3/.../9/free）
//   - variant: 布局变体 ID（可省略取第一个）
//   - rows / cols: 自由拼图时使用
//   - cell: 单格边长 px（默认 400）
//   - gap: 间距 px（默认 10）
//   - bg: 背景色 #RRGGBB（默认 #ffffff）
//   - cells: 每格视图调整 JSON 数组（与 images 顺序对应）[{"zoom":1.2,"dx":10,"dy":-5}]
//   - format: 输出格式 jpg/png（默认 jpg）
//   - quality: jpg 质量（默认 90）
func handleCollage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		jsonError(w, 405, "仅支持 POST")
		return
	}
	imgs, _, err := parseImages(r, "images")
	if err != nil {
		jsonError(w, 400, err.Error())
		return
	}
	mode := r.FormValue("mode")
	if mode == "" {
		mode = "classic9"
	}
	variant := r.FormValue("variant")
	cellSize := atoiDefault(r.FormValue("cell"), 400)
	gap := atoiDefault(r.FormValue("gap"), 10)
	format := strings.ToLower(r.FormValue("format"))
	if format == "" {
		format = "jpg"
	}
	quality := atoiDefault(r.FormValue("quality"), 90)

	// 每格视图调整
	var views []CellView
	if cellsStr := r.FormValue("cells"); cellsStr != "" {
		if err := json.Unmarshal([]byte(cellsStr), &views); err != nil {
			jsonError(w, 400, "cells 参数解析失败: "+err.Error())
			return
		}
	}

	if cellSize < 100 || cellSize > 2000 {
		cellSize = 400
	}
	if gap < 0 || gap > 200 {
		gap = 10
	}

	br, bg_, bb, _ := parseHexColor(r.FormValue("bg"))
	bgColor := color.NRGBA{R: br, G: bg_, B: bb, A: 255}

	var layout *Layout
	if mode == "free" {
		rows := atoiDefault(r.FormValue("rows"), 2)
		cols := atoiDefault(r.FormValue("cols"), 2)
		if rows < 1 || rows > 7 || cols < 1 || cols > 7 {
			jsonError(w, 400, "自由拼图行/列数需在 1-7 之间")
			return
		}
		layout = &Layout{ID: "free", Name: "自由拼图", Rows: rows, Cols: cols, Cells: gridCells(rows, cols, true)}
	} else {
		layout = findCollageLayout(mode, variant)
		if layout == nil {
			jsonError(w, 400, "未知的模板模式: "+mode)
			return
		}
	}

	need := layoutImageCount(layout)
	if len(imgs) < need {
		jsonError(w, 400, fmt.Sprintf("模板 %s 需要 %d 张图片，当前只上传了 %d 张", layout.Name, need, len(imgs)))
		return
	}

	out, err := buildCollage(layout, imgs[:need], cellSize, gap, bgColor, views)
	if err != nil {
		jsonError(w, 500, "拼图失败: "+err.Error())
		return
	}
	sendImage(w, out, format, "collage."+extFor(format), quality)
}

func extFor(format string) string {
	switch strings.ToLower(format) {
	case "png":
		return "png"
	case "bmp":
		return "bmp"
	case "gif":
		return "gif"
	case "tiff":
		return "tiff"
	default:
		return "jpg"
	}
}

// buildCollage 依据布局把图片拼成一张图（支持每格缩放/平移）
func buildCollage(l *Layout, imgs []image.Image, cellSize, gap int, bg color.Color, views []CellView) (*image.NRGBA, error) {
	w := l.Cols*cellSize + gap*(l.Cols-1)
	h := l.Rows*cellSize + gap*(l.Rows-1)
	canvas := imaging.New(w, h, bg)

	cellW := cellSize
	cellH := cellSize
	imgIdx := 0
	for _, c := range l.Cells {
		x := c.Col * (cellW + gap)
		y := c.Row * (cellH + gap)
		cw := c.ColSpan*cellW + (c.ColSpan-1)*gap
		ch := c.RowSpan*cellH + (c.RowSpan-1)*gap
		if !c.HasImage {
			continue
		}
		if imgIdx >= len(imgs) {
			break
		}
		v := CellView{Zoom: 1}
		if imgIdx < len(views) {
			v = views[imgIdx]
		}
		filled := fillCoverView(imgs[imgIdx], cw, ch, v)
		canvas = imaging.Paste(canvas, filled, image.Pt(x, y))
		imgIdx++
	}
	return canvas, nil
}

// fillCoverView 按视图参数（缩放/平移）把图片裁剪填充到目标尺寸
// 算法：先将图片 cover 缩放得到"世界"（世界尺寸至少等于目标格子），
// 再按 zoom 决定可见窗口大小，按 dx/dy（占窗口百分比）移动窗口，裁剪后缩放回格子尺寸。
func fillCoverView(src image.Image, w, h int, v CellView) *image.NRGBA {
	b := src.Bounds()
	srcW, srcH := b.Dx(), b.Dy()
	if srcW <= 0 || srcH <= 0 {
		return imaging.New(w, h, color.NRGBA{0, 0, 0, 255})
	}
	zoom := v.Zoom
	if zoom <= 0 {
		zoom = 1
	}
	zoom = clampFloat(zoom, 0.5, 5)

	s0 := math.Max(float64(w)/float64(srcW), float64(h)/float64(srcH))
	sw := int(math.Round(float64(srcW) * s0))
	sh := int(math.Round(float64(srcH) * s0))
	world := imaging.Resize(src, sw, sh, imaging.Lanczos)

	ww, wh := float64(w), float64(h)
	worldW, worldH := float64(sw), float64(sh)
	if zoom < 1 {
		// 缩小视野：放大世界，窗口保持格子大小
		f := 1.0 / zoom
		worldW = float64(sw) * f
		worldH = float64(sh) * f
		world = imaging.Resize(world, int(worldW), int(worldH), imaging.Lanczos)
	} else {
		// 放大视野：窗口缩小
		ww = float64(w) / zoom
		wh = float64(h) / zoom
	}

	cx := worldW/2 + (v.DX/100)*ww
	cy := worldH/2 + (v.DY/100)*wh
	// 夹紧窗口中心，防止越界
	maxCX := worldW - ww/2
	maxCY := worldH - wh/2
	if ww > worldW {
		cx = worldW / 2
	} else {
		cx = clampFloat(cx, ww/2, maxCX)
	}
	if wh > worldH {
		cy = worldH / 2
	} else {
		cy = clampFloat(cy, wh/2, maxCY)
	}

	x0 := int(math.Round(cx - ww/2))
	y0 := int(math.Round(cy - wh/2))
	x1 := int(math.Round(cx + ww/2))
	y1 := int(math.Round(cy + wh/2))
	if x0 < 0 {
		x0 = 0
	}
	if y0 < 0 {
		y0 = 0
	}
	if x1 > int(worldW) {
		x1 = int(worldW)
	}
	if y1 > int(worldH) {
		y1 = int(worldH)
	}
	if x1 <= x0 || y1 <= y0 {
		return imaging.New(w, h, color.NRGBA{0, 0, 0, 255})
	}
	cropped := imaging.Crop(world, image.Rect(x0, y0, x1, y1))
	return imaging.Resize(cropped, w, h, imaging.Lanczos)
}

func atoiDefault(s string, def int) int {
	if s == "" {
		return def
	}
	v, err := strconv.Atoi(s)
	if err != nil {
		return def
	}
	return v
}
