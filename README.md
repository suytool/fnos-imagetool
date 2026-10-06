# 图片工具 (Image Tool) — 飞牛 fnOS 应用

飞牛 fnOS 第三方应用（FPK），Go 编写，单静态二进制，服务端口 **3434**。
图片处理全部在本地 NAS 完成，不上传任何外部服务器。

## 功能

| 模块 | 说明 |
|---|---|
| 拼图拆分 | 经典九宫格 / 2~9 图模板 / 自由拼图 7×7；**格子内图片支持缩放、拖动微调**；分图 1拆9/1拆6/1拆4、自由分图 10×10 |
| 添加水印 | 文字水印（字号/颜色/透明度/位置/旋转，内置思源黑体中文字体）+ 图片水印 |
| 图片修改 | 缩放 / 旋转 / 比例裁剪 / 灰度 / 反色 / 亮度对比度饱和度 / 格式转换（JPG/PNG/BMP/GIF/TIFF） |
| 支持作者 | 顶栏「❤ 支持作者」弹窗展示支付宝/微信收款码 |

## 安装

1. 到 [Releases](https://github.com/suytool/fnos-imagetool/releases) 下载对应架构的安装包：
   - `图片工具-0.0.2-x86_64.fpk` — Intel / AMD 平台
   - `图片工具-0.0.2-arm64.fpk` — ARM 平台（鲲鹏 / 飞腾 / 树莓派等）
2. 飞牛应用中心 → 手动安装 → 上传 .fpk → 确认（第三方应用提示）→ 桌面出现「图片工具」图标。
3. 点击图标在浏览器打开，端口 3434。

## 目录结构

```
imagetool-src/           Go 源码（编译产物即运行二进制）
├── main.go              路由与静态服务
├── collage.go           拼图（含每格缩放/平移视图渲染）
├── split.go             分图（zip 打包）
├── watermark.go         文字/图片水印
├── edit.go              图片修改（ops JSON 驱动）
├── layout.go            拼图/分图模板定义（前端 COLLAGE 对象与此同步）
├── font.go              中文字体渲染（思源黑体 OTF）
├── imageutil.go         通用工具
└── web/                 前端资源（index.html / style.css / app.js / font.otf / alipay.png / wechat.png）

imagetool-app/           FPK 打包工程（fnpack build -d . 用）
├── manifest             FPK 元信息（appname=com.imagetool.app, version=0.0.2, platform=x86|arm）
├── cmd/                 生命周期脚本（main: start/stop/status；其余为钩子）
├── config/              privilege 与 resource 声明
├── ICON.PNG / ICON_256.PNG
└── app/                 server/（二进制+web）、ui/（桌面入口+图标）
```

## 构建

Go ≥ 1.26（依赖 golang.org/x/image v0.46.0）。

```bash
cd imagetool-src

# 本地运行
go build -o imagetool .
./imagetool        # 默认 0.0.0.0:3434，可用环境变量 PORT 覆盖

# 交叉编译静态二进制
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o imagetool-amd64 .
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags "-s -w" -o imagetool-arm64 .
```

注意：本项目不使用 `go:embed`，静态资源运行时按 `cwd/web → exeDir/web → exeDir/../www → exeDir/www` 查找。

## 打包 FPK

fnpack（官方 v1.2.3）：`https://static2.fnnas.com/fnpack/fnpack-1.2.3-linux-amd64`

```bash
cp imagetool-amd64 imagetool-app/app/server/imagetool   # manifest platform=x86
cd imagetool-app && /tmp/fnpack build -d .              # 产出 com.imagetool.app.fpk
# arm64：manifest 改 platform=arm，放 arm64 二进制，再 build
```

## FPK 生命周期契约

`cmd/main`：`start` 启动服务、`stop` 停止、`status` 退出码 0=运行中 / 3=未运行；
PID 记录于 `$TRIM_PKGVAR/app.pid`，日志在 `$TRIM_PKGVAR/info.log`。

## 接口

`POST /api/collage`（拼图，支持 cells 每格缩放/平移）、`POST /api/split`（分图 zip）、
`POST /api/watermark`、`POST /api/edit`、`GET /api/health`。
