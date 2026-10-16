package main

import (
	"bytes"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"strings"

	"github.com/disintegration/imaging"
	"golang.org/x/image/bmp"
	"golang.org/x/image/tiff"
)

// decodeImage 自动识别格式解码图片
func decodeImage(data []byte) (image.Image, string, error) {
	img, format, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		return nil, "", err
	}
	return img, format, nil
}

// encodeImage 按格式编码图片
func encodeImage(img image.Image, format string, quality int) ([]byte, error) {
	var buf bytes.Buffer
	switch strings.ToLower(format) {
	case "png":
		err := png.Encode(&buf, img)
		return buf.Bytes(), err
	case "bmp":
		err := bmp.Encode(&buf, img)
		return buf.Bytes(), err
	case "gif":
		err := gif.Encode(&buf, img, nil)
		return buf.Bytes(), err
	case "tiff":
		err := tiff.Encode(&buf, img, nil)
		return buf.Bytes(), err
	case "jpg", "jpeg":
		if quality <= 0 {
			quality = 90
		}
		if quality > 100 {
			quality = 100
		}
		err := jpeg.Encode(&buf, img, &jpeg.Options{Quality: quality})
		return buf.Bytes(), err
	default:
		return nil, fmt.Errorf("不支持的输出格式: %s", format)
	}
}

// fillCover 将图片按 cover 模式裁剪填充到目标尺寸（居中裁剪）
func fillCover(src image.Image, w, h int) *image.NRGBA {
	return imaging.Fill(src, w, h, imaging.Center, imaging.Lanczos)
}

// clampInt 限制整数范围
func clampInt(v, lo, hi int) int {
	if v < lo {
		return lo
	}
	if v > hi {
		return hi
	}
	return v
}

func clampFloat(v, lo, hi float64) float64 {
	if v < lo {
		return lo
	}
	if v > hi {
		return hi
	}
	return v
}

// parseHexColor 解析 #RRGGBB / #RRGGBBAA
func parseHexColor(s string) (r, g, b, a uint8) {
	s = strings.TrimPrefix(strings.TrimSpace(s), "#")
	a = 255
	switch len(s) {
	case 3:
		r = hexPair(s[0])
		g = hexPair(s[1])
		b = hexPair(s[2])
	case 6:
		r = hexPair(s[0])<<4 | hexPair(s[1])
		g = hexPair(s[2])<<4 | hexPair(s[3])
		b = hexPair(s[4])<<4 | hexPair(s[5])
	case 8:
		r = hexPair(s[0])<<4 | hexPair(s[1])
		g = hexPair(s[2])<<4 | hexPair(s[3])
		b = hexPair(s[4])<<4 | hexPair(s[5])
		a = hexPair(s[6])<<4 | hexPair(s[7])
	}
	return
}

func hexPair(c byte) uint8 {
	switch {
	case c >= '0' && c <= '9':
		return c - '0'
	case c >= 'a' && c <= 'f':
		return c - 'a' + 10
	case c >= 'A' && c <= 'F':
		return c - 'A' + 10
	}
	return 0
}
