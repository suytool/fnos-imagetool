package main

import (
	"image"
	"image/color"
	"image/draw"
	"math"
	"net/http"
	"strings"

	"github.com/disintegration/imaging"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

// handleWatermark /api/watermark
// 表单字段：
//   - image: 原图
//   - text: 水印文字（可选）
//   - font_size: 字号 px（默认 48）
//   - color: 文字颜色 #RRGGBB（默认 #ffffff）
//   - opacity: 不透明度 0-100（默认 100）
//   - position: tl/tr/bl/br/center/tiled（默认 br）
//   - rotation: 旋转角度（默认 0）
//   - logo: 图片水印文件（可选）
//   - logo_size: 水印图宽 px（默认 200，按比例缩放）
//   - logo_opacity: 水印图不透明度 0-100（默认 100）
//   - format: jpg/png
//   - quality: jpg 质量
func handleWatermark(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		jsonError(w, 405, "仅支持 POST")
		return
	}
	img, _, err := parseSingleImage(r, "image")
	if err != nil {
		jsonError(w, 400, err.Error())
		return
	}
	text := strings.TrimSpace(r.FormValue("text"))
	position := strings.ToLower(r.FormValue("position"))
	if position == "" {
		position = "br"
	}
	fontSize := atoiDefault(r.FormValue("font_size"), 48)
	opacity := atoiDefault(r.FormValue("opacity"), 100)
	rotation := atoiDefault(r.FormValue("rotation"), 0)
	cr, cg, cb, ca := parseHexColor(r.FormValue("color"))
	if ca == 0 {
		ca = 255
	}
	format := strings.ToLower(r.FormValue("format"))
	if format == "" {
		format = "jpg"
	}
	quality := atoiDefault(r.FormValue("quality"), 90)

	opacityF := clampFloat(float64(opacity)/100.0, 0, 1)

	out := imaging.Clone(img)

	// 文字水印
	if text != "" {
		if fontSize < 8 || fontSize > 2000 {
			fontSize = 48
		}
		textImg, err := renderTextWatermark(text, fontSize, color.NRGBA{R: cr, G: cg, B: cb, A: 255}, opacityF, rotation)
		if err != nil {
			jsonError(w, 500, "文字水印渲染失败: "+err.Error())
			return
		}
		out = pasteWatermark(out, textImg, position, 20)
	}

	// 图片水印
	if fhs, ok := r.MultipartForm.File["logo"]; ok && len(fhs) > 0 {
		data, err := readUploadedFile(fhs[0])
		if err != nil {
			jsonError(w, 400, "读取水印图片失败: "+err.Error())
			return
		}
		logo, _, err := decodeImage(data)
		if err != nil {
			jsonError(w, 400, "水印图片无法解析: "+err.Error())
			return
		}
		logoW := atoiDefault(r.FormValue("logo_size"), 200)
		logoOpacity := atoiDefault(r.FormValue("logo_opacity"), 100)
		logoOpacityF := clampFloat(float64(logoOpacity)/100.0, 0, 1)
		logoImg := scaleToWidth(logo, logoW)
		logoImg = applyOpacity(logoImg, logoOpacityF)
		out = pasteWatermark(out, logoImg, position, 20)
	}

	sendImage(w, out, format, "watermarked."+extFor(format), quality)
}

// scaleToWidth 等比缩放到指定宽度
func scaleToWidth(src image.Image, w int) *image.NRGBA {
	if w <= 0 {
		w = 200
	}
	b := src.Bounds()
	if b.Dx() <= 0 {
		return imaging.Clone(src)
	}
	ratio := float64(w) / float64(b.Dx())
	h := int(math.Round(float64(b.Dy()) * ratio))
	return imaging.Resize(src, w, h, imaging.Lanczos)
}

// applyOpacity 对整个图片应用不透明度（生成带 alpha 的 NRGBA）
func applyOpacity(src image.Image, opacity float64) *image.NRGBA {
	b := src.Bounds()
	dst := image.NewNRGBA(image.Rect(0, 0, b.Dx(), b.Dy()))
	a := uint8(math.Round(opacity * 255))
	for y := 0; y < b.Dy(); y++ {
		for x := 0; x < b.Dx(); x++ {
			c := color.NRGBAModel.Convert(src.At(b.Min.X+x, b.Min.Y+y)).(color.NRGBA)
			// 保留原图自身 alpha，再乘水印不透明度
			c.A = uint8(int(c.A) * int(a) / 255)
			dst.SetNRGBA(x, y, c)
		}
	}
	return dst
}

// pasteWatermark 将水印图片混合（Over）到目标图片的指定位置
func pasteWatermark(dst *image.NRGBA, wm image.Image, position string, margin int) *image.NRGBA {
	db := dst.Bounds()
	wb := wm.Bounds()
	wmw, wmh := wb.Dx(), wb.Dy()
	var x, y int
	switch position {
	case "tl":
		x, y = margin, margin
	case "tr":
		x, y = db.Dx()-wmw-margin, margin
	case "bl":
		x, y = margin, db.Dy()-wmh-margin
	case "center":
		x, y = (db.Dx()-wmw)/2, (db.Dy()-wmh)/2
	case "tiled":
		// 平铺
		for ty := margin; ty+wmh <= db.Dy()-margin; ty += wmh + margin {
			for tx := margin; tx+wmw <= db.Dx()-margin; tx += wmw + margin {
				draw.Draw(dst, image.Rect(tx, ty, tx+wmw, ty+wmh), wm, wb.Min, draw.Over)
			}
		}
		return dst
	default: // br
		x, y = db.Dx()-wmw-margin, db.Dy()-wmh-margin
	}
	draw.Draw(dst, image.Rect(x, y, x+wmw, y+wmh), wm, wb.Min, draw.Over)
	return dst
}

// renderTextWatermark 渲染文字水印图层（支持旋转）
func renderTextWatermark(text string, fontSize int, textColor color.NRGBA, opacity float64, rotation int) (*image.NRGBA, error) {
	face, err := newFace(float64(fontSize))
	if err != nil {
		return nil, err
	}
	m := face.Metrics()
	bounds, _ := font.BoundString(face, text)
	w := bounds.Max.X.Ceil() - bounds.Min.X.Ceil()
	h := (m.Ascent + m.Descent).Ceil()
	if w < 1 {
		w = 1
	}
	if h < 1 {
		h = 1
	}
	pad := 8
	canvasW := w + pad*2
	canvasH := h + pad*2
	rgba := image.NewRGBA(image.Rect(0, 0, canvasW, canvasH))
	d := &font.Drawer{
		Dst:  rgba,
		Src:  image.NewUniform(textColor),
		Face: face,
		Dot:  fixed.P(pad, pad+m.Ascent.Ceil()),
	}
	d.DrawString(text)

	// 旋转
	var rot *image.NRGBA
	if rotation%360 != 0 {
		rot = imaging.Rotate(rgba, float64(rotation), color.NRGBA{0, 0, 0, 0})
	} else {
		rot = imaging.Clone(rgba)
	}
	return applyOpacity(rot, opacity), nil
}
