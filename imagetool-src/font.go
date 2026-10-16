package main

import (
	"os"
	"path/filepath"
	"sync"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

var (
	fontOnce sync.Once
	parsed   *opentype.Font
	fontErr  error
)

// loadFontData 从资源目录读取内置字体（思源黑体）
func loadFontData() ([]byte, error) {
	return readAsset("font.otf")
}

// embeddedFont 解析内置思源黑体
func embeddedFont() (*opentype.Font, error) {
	fontOnce.Do(func() {
		data, err := loadFontData()
		if err != nil {
			fontErr = err
			return
		}
		parsed, fontErr = opentype.Parse(data)
	})
	return parsed, fontErr
}

// newFace 创建指定像素字号的字体面
func newFace(px float64) (font.Face, error) {
	f, err := embeddedFont()
	if err != nil {
		return nil, err
	}
	return opentype.NewFace(f, &opentype.FaceOptions{
		Size:    px,
		DPI:     72,
		Hinting: font.HintingFull,
	})
}

// readAsset 从多个候选目录查找资源文件（web 静态资源/字体）
func readAsset(name string) ([]byte, error) {
	for _, dir := range assetDirs() {
		p := filepath.Join(dir, name)
		if data, err := os.ReadFile(p); err == nil {
			return data, nil
		}
	}
	return nil, os.ErrNotExist
}

// assetDirs 候选资源目录：
//  1. 当前工作目录 web/
//  2. 可执行文件所在目录 web/
//  3. 可执行文件所在目录 ../www/（fnpack 模板布局）
//  4. 可执行文件所在目录 www/
func assetDirs() []string {
	exe, err := os.Executable()
	if err != nil {
		exe = "."
	}
	exeDir := filepath.Dir(exe)
	cwd, _ := os.Getwd()
	seen := map[string]bool{}
	var dirs []string
	add := func(d string) {
		if d != "" && !seen[d] {
			seen[d] = true
			dirs = append(dirs, d)
		}
	}
	add(filepath.Join(cwd, "web"))
	add(filepath.Join(exeDir, "web"))
	add(filepath.Join(exeDir, "..", "www"))
	add(filepath.Join(exeDir, "www"))
	return dirs
}
