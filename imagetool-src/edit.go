package main

import (
	"encoding/json"
	"image"
	"image/color"
	"net/http"
	"strings"

	"github.com/disintegration/imaging"
)

// EditOps 图片修改操作参数
type EditOps struct {
	Width      int     `json:"width"`      // 目标宽（resize 使用）
	Height     int     `json:"height"`     // 目标高
	KeepRatio  bool    `json:"keep_ratio"` // 等比缩放
	Scale      float64 `json:"scale"`      // 按比例缩放（0 表示不用）
	Rotate     float64 `json:"rotate"`     // 旋转角度
	CropX      int     `json:"crop_x"`
	CropY      int     `json:"crop_y"`
	CropW      int     `json:"crop_w"`
	CropH      int     `json:"crop_h"`
	CropAspect string  `json:"crop_aspect"` // 裁剪比例：1:1 / 4:3 / 3:4 / 16:9 / 9:16 / 3:2
	Grayscale  bool    `json:"grayscale"`
	Invert     bool    `json:"invert"`
	Brightness int     `json:"brightness"` // -100 ~ 100
	Contrast   int     `json:"contrast"`   // -100 ~ 100
	Saturation int     `json:"saturation"` // -100 ~ 100
	Format     string  `json:"format"`     // jpg/png/bmp/gif/tiff
	Quality    int     `json:"quality"`    // jpg 质量
}

// handleEdit /api/edit
// 表单字段：image（单图）、ops（JSON 字符串）
func handleEdit(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		jsonError(w, 405, "仅支持 POST")
		return
	}
	img, _, err := parseSingleImage(r, "image")
	if err != nil {
		jsonError(w, 400, err.Error())
		return
	}
	var ops EditOps
	opsStr := r.FormValue("ops")
	if opsStr == "" {
		jsonError(w, 400, "缺少 ops 参数")
		return
	}
	if err := json.Unmarshal([]byte(opsStr), &ops); err != nil {
		jsonError(w, 400, "ops 参数解析失败: "+err.Error())
		return
	}
	if ops.Format == "" {
		ops.Format = "jpg"
	}
	if ops.Quality <= 0 {
		ops.Quality = 90
	}

	out, err := applyEdits(img, ops)
	if err != nil {
		jsonError(w, 400, "修改失败: "+err.Error())
		return
	}
	sendImage(w, out, ops.Format, "edited."+extFor(ops.Format), ops.Quality)
}

func applyEdits(src image.Image, ops EditOps) (*image.NRGBA, error) {
	b := src.Bounds()
	W, H := b.Dx(), b.Dy()
	img := imaging.Clone(src)

	// 1. 按比例缩放
	if ops.Scale > 0 && ops.Scale != 1 {
		img = imaging.Resize(img, int(float64(W)*ops.Scale), int(float64(H)*ops.Scale), imaging.Lanczos)
	}

	// 2. 指定尺寸缩放
	if ops.Width > 0 || ops.Height > 0 {
		if ops.KeepRatio || (ops.Width > 0 && ops.Height == 0) || (ops.Width == 0 && ops.Height > 0) {
			img = imaging.Resize(img, ops.Width, ops.Height, imaging.Lanczos)
		} else {
			img = imaging.Resize(img, ops.Width, ops.Height, imaging.Lanczos)
		}
	}

	// 3. 旋转
	if ops.Rotate != 0 {
		img = imaging.Rotate(img, ops.Rotate, color.NRGBA{0, 0, 0, 0})
	}

	// 4. 裁剪
	cb := img.Bounds()
	if ops.CropW > 0 && ops.CropH > 0 {
		x0 := clampInt(ops.CropX, 0, cb.Dx()-1)
		y0 := clampInt(ops.CropY, 0, cb.Dy()-1)
		cw := clampInt(ops.CropW, 1, cb.Dx()-x0)
		ch := clampInt(ops.CropH, 1, cb.Dy()-y0)
		img = imaging.Crop(img, image.Rect(x0, y0, x0+cw, y0+ch))
	} else if ops.CropAspect != "" {
		cw, ch := aspectSize(cb.Dx(), cb.Dy(), ops.CropAspect)
		if cw > 0 && ch > 0 {
			x0 := (cb.Dx() - cw) / 2
			y0 := (cb.Dy() - ch) / 2
			img = imaging.Crop(img, image.Rect(x0, y0, x0+cw, y0+ch))
		}
	}

	// 5. 颜色调整
	if ops.Grayscale {
		img = imaging.Grayscale(img)
	}
	if ops.Invert {
		img = imaging.Invert(img)
	}
	if ops.Brightness != 0 {
		img = imaging.AdjustBrightness(img, float64(clampInt(ops.Brightness, -100, 100)))
	}
	if ops.Contrast != 0 {
		img = imaging.AdjustContrast(img, float64(clampInt(ops.Contrast, -100, 100)))
	}
	if ops.Saturation != 0 {
		img = imaging.AdjustSaturation(img, float64(clampInt(ops.Saturation, -100, 100)))
	}
	return img, nil
}

// aspectSize 根据目标比例计算裁剪尺寸（居中裁剪）
func aspectSize(w, h int, aspect string) (int, int) {
	var aw, ah int
	switch strings.ToLower(strings.ReplaceAll(aspect, " ", "")) {
	case "1:1":
		aw, ah = 1, 1
	case "4:3":
		aw, ah = 4, 3
	case "3:4":
		aw, ah = 3, 4
	case "16:9":
		aw, ah = 16, 9
	case "9:16":
		aw, ah = 9, 16
	case "3:2":
		aw, ah = 3, 2
	case "2:3":
		aw, ah = 2, 3
	case "21:9":
		aw, ah = 21, 9
	default:
		return 0, 0
	}
	if float64(w)/float64(h) > float64(aw)/float64(ah) {
		// 图更宽：以高度为基准
		cw := h * aw / ah
		return cw, h
	}
	ch := w * ah / aw
	return w, ch
}
