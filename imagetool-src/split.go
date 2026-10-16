package main

import (
	"archive/zip"
	"bytes"
	"fmt"
	"image"
	"net/http"
	"strings"

	"github.com/disintegration/imaging"
)

// handleSplit /api/split
// 表单字段：
//   - image: 单张图片
//   - rows / cols: 拆分行列（1-10）
//   - format: 分片格式 jpg/png（默认 jpg）
//   - quality: jpg 质量（默认 90）
// 返回 zip 包，内含按顺序命名的分片
func handleSplit(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		jsonError(w, 405, "仅支持 POST")
		return
	}
	img, _, err := parseSingleImage(r, "image")
	if err != nil {
		jsonError(w, 400, err.Error())
		return
	}
	rows := atoiDefault(r.FormValue("rows"), 3)
	cols := atoiDefault(r.FormValue("cols"), 3)
	if rows < 1 || rows > 10 || cols < 1 || cols > 10 {
		jsonError(w, 400, "拆分行列数需在 1-10 之间")
		return
	}
	format := strings.ToLower(r.FormValue("format"))
	if format == "" {
		format = "jpg"
	}
	quality := atoiDefault(r.FormValue("quality"), 90)

	zipData, err := buildSplitZip(img, rows, cols, format, quality)
	if err != nil {
		jsonError(w, 500, "分图失败: "+err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/zip")
	w.Header().Set("Content-Disposition", fmt.Sprintf(`attachment; filename="split_%dx%d.zip"`, rows, cols))
	w.Header().Set("Content-Length", fmt.Sprint(len(zipData)))
	w.Write(zipData)
}

// buildSplitZip 把图片均分为 rows×cols 并打包 zip
func buildSplitZip(src image.Image, rows, cols int, format string, quality int) ([]byte, error) {
	b := src.Bounds()
	W := b.Dx()
	H := b.Dy()
	pieceW := W / cols
	pieceH := H / rows
	if pieceW < 1 || pieceH < 1 {
		return nil, fmt.Errorf("图片尺寸过小，无法拆分为 %d×%d", rows, cols)
	}
	var buf bytes.Buffer
	zw := zip.NewWriter(&buf)
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			// 边缘格子取剩余像素
			x0 := col * pieceW
			y0 := row * pieceH
			x1 := x0 + pieceW
			y1 := y0 + pieceH
			if col == cols-1 {
				x1 = W
			}
			if row == rows-1 {
				y1 = H
			}
			piece := imaging.Crop(src, image.Rect(x0, y0, x1, y1))
			data, err := encodeImage(piece, format, quality)
			if err != nil {
				zw.Close()
				return nil, err
			}
			name := fmt.Sprintf("part_%02d_%02d.%s", row+1, col+1, extFor(format))
			fw, err := zw.Create(name)
			if err != nil {
				zw.Close()
				return nil, err
			}
			if _, err := fw.Write(data); err != nil {
				zw.Close()
				return nil, err
			}
		}
	}
	if err := zw.Close(); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
