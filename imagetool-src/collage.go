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
	Zoom float64 `json:"zoom"` // 缩放倍数 0.01-5（1%-500%），1 表示 contain 完整显示
	DX   float64 `json:"dx"`   // 水平偏移百分比（相对格子宽度），仅图片超出格子时生效
	DY   float64 `json:"dy"`   // 垂直偏移百分比（相对格子高度）
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
		filled := fillContainView(imgs[imgIdx], cw, ch, v, bg)
		canvas = imaging.Paste(canvas, filled, image.Pt(x, y))
		imgIdx++
	}
	return canvas, nil
}

// fillContainView 按视图参数（缩放/平移）把图片渲染到目标格子。
// 100% 时图片按 contain 完整显示在格内（不强制填满）；
// <100% 缩小并居中（图片周围为背景色）；>100% 放大，超出格子的方向按窗口裁剪并支持平移。
// dx/dy 为占格子宽高的百分比。
func fillContainView(src image.Image, w, h int, v CellView, bg color.Color) *image.NRGBA {
	b := src.Bounds()
	srcW, srcH := b.Dx(), b.Dy()
	if srcW <= 0 || srcH <= 0 {
		return imaging.New(w, h, bg)
	}
	zoom := v.Zoom
	if zoom <= 0 {
		zoom = 1
	}
	zoom = clampFloat(zoom, 0.01, 5)

	s0 := math.Min(float64(w)/float64(srcW), float64(h)/float64(srcH))
	effW := float64(srcW) * s0 * zoom
	effH := float64(srcH) * s0 * zoom
	canvas := imaging.New(w, h, bg)

	if effW <= float64(w) && effH <= float64(h) {
		// 图片小于（等于）格子：完整显示，居中，不裁剪
		iw := int(math.Round(effW))
		ih := int(math.Round(effH))
		if iw <= 0 || ih <= 0 {
			return canvas
		}
		resized := imaging.Resize(src, iw, ih, imaging.Lanczos)
		x0 := (w - iw) / 2
		y0 := (h - ih) / 2
		return imaging.Paste(canvas, resized, image.Pt(x0, y0))
	}

	// 图片超过格子：放大到 effW×effH，再按窗口（格子大小）裁剪，支持平移
	iw := int(math.Round(effW))
	ih := int(math.Round(effH))
	world := imaging.Resize(src, iw, ih, imaging.Lanczos)
	cx := effW/2 + (v.DX/100)*float64(w)
	cy := effH/2 + (v.DY/100)*float64(h)
	if effW > float64(w) {
		cx = clampFloat(cx, float64(w)/2, effW-float64(w)/2)
	} else {
		cx = effW / 2
	}
	if effH > float64(h) {
		cy = clampFloat(cy, float64(h)/2, effH-float64(h)/2)
	} else {
		cy = effH / 2
	}
	x0 := int(math.Round(cx - float64(w)/2))
	y0 := int(math.Round(cy - float64(h)/2))
	x1 := x0 + w
	y1 := y0 + h
	if x0 < 0 {
		x0 = 0
	}
	if y0 < 0 {
		y0 = 0
	}
	if x1 > iw {
		x1 = iw
	}
	if y1 > ih {
		y1 = ih
	}
	if x1 <= x0 || y1 <= y0 {
		return canvas
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
