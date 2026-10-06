// 图片工具 (Image Tool) - fnOS FPK 应用后端
// 版本 0.0.1，监听端口 3434
// 功能：拼图、分图、添加水印、图片修改
package main

import (
	"encoding/json"
	"fmt"
	"image"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

const maxUploadSize = 100 << 20 // 100MB

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3434"
	}
	addr := "0.0.0.0:" + port

	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/api/collage", handleCollage)
	http.HandleFunc("/api/split", handleSplit)
	http.HandleFunc("/api/watermark", handleWatermark)
	http.HandleFunc("/api/edit", handleEdit)
	http.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})

	log.Printf("ImageTool 0.0.1 listening on %s", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	name := strings.TrimPrefix(r.URL.Path, "/")
	if name == "" {
		name = "index.html"
	}
	if strings.Contains(name, "..") || strings.Contains(name, "\\") {
		http.NotFound(w, r)
		return
	}
	data, err := readAsset(name)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	ct := contentTypeByExt(name)
	w.Header().Set("Content-Type", ct)
	w.Header().Set("Cache-Control", "no-cache")
	w.Write(data)
}

func contentTypeByExt(name string) string {
	switch strings.ToLower(filepath.Ext(name)) {
	case ".html":
		return "text/html; charset=utf-8"
	case ".css":
		return "text/css; charset=utf-8"
	case ".js":
		return "application/javascript; charset=utf-8"
	case ".png":
		return "image/png"
	case ".jpg", ".jpeg":
		return "image/jpeg"
	case ".svg":
		return "image/svg+xml"
	case ".woff", ".woff2":
		return "font/woff2"
	case ".ttf":
		return "font/ttf"
	case ".otf":
		return "font/otf"
	case ".ico":
		return "image/x-icon"
	default:
		return "application/octet-stream"
	}
}

// 读取 multipart 中的单个文件字段
func readUploadedFile(fh *multipart.FileHeader) ([]byte, error) {
	f, err := fh.Open()
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return io.ReadAll(io.LimitReader(f, maxUploadSize))
}

// 从请求中解析一张或多张图片
func parseImages(r *http.Request, field string) ([]image.Image, []string, error) {
	if err := r.ParseMultipartForm(maxUploadSize); err != nil {
		return nil, nil, err
	}
	var imgs []image.Image
	var names []string
	fhs := r.MultipartForm.File[field]
	if len(fhs) == 0 {
		return nil, nil, fmt.Errorf("未找到图片文件字段: %s", field)
	}
	for _, fh := range fhs {
		data, err := readUploadedFile(fh)
		if err != nil {
			return nil, nil, err
		}
		img, _, err := decodeImage(data)
		if err != nil {
			return nil, nil, fmt.Errorf("无法解析图片 %s: %v", fh.Filename, err)
		}
		imgs = append(imgs, img)
		names = append(names, fh.Filename)
	}
	return imgs, names, nil
}

func parseSingleImage(r *http.Request, field string) (image.Image, string, error) {
	imgs, names, err := parseImages(r, field)
	if err != nil {
		return nil, "", err
	}
	return imgs[0], names[0], nil
}

func jsonError(w http.ResponseWriter, status int, msg string) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]string{"error": msg})
}

// sendImage 输出处理后的图片
func sendImage(w http.ResponseWriter, img image.Image, format, filename string, quality int) {
	data, err := encodeImage(img, format, quality)
	if err != nil {
		jsonError(w, 500, "图片编码失败: "+err.Error())
		return
	}
	ct := "image/jpeg"
	switch format {
	case "png":
		ct = "image/png"
	case "bmp":
		ct = "image/bmp"
	case "gif":
		ct = "image/gif"
	case "tiff":
		ct = "image/tiff"
	}
	w.Header().Set("Content-Type", ct)
	w.Header().Set("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, filename))
	w.Header().Set("Content-Length", fmt.Sprint(len(data)))
	w.Write(data)
}
