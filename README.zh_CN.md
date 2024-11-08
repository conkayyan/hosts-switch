## README.md

- en [English](README.md)

## 关于

这是一个基于 [wails](https://wails.io/) 开发的简单桌面应用程序，兼容 Linux、Windows 和 Mac，可以帮助您更好地维护本地主机配置。

## 屏幕截图

![Screenshot1](screenshot1.png "Screenshot1")
![Screenshot2](screenshot2.png "Screenshot2")

## 环境要求
基于 go v1.21+、wails v2.8.2、nodejs v20。

在 Windows 构建的 Wails 应用程序运行时依赖 Microsoft WebView2 Runtime 。Windows 11 默认安装 WebView2 ，但有些系统默认不会安装。

## 安装 Wails
执行 `go install github.com/wailsapp/wails/v2/cmd/wails@latest` 来安装 Wails CLI。

## 实时开发

要在实时开发模式下运行，请在项目目录中运行 `wails dev`。这将运行一个 Vite 开发
服务器，它将提供非常快速的前端更改热重载。如果您想在浏览器中开发
并可以访问您的 Go 方法，还有一个在 http://localhost:34115 上运行的开发服务器。在您的浏览器中连接到此服务器，您就可以从 devtools 调用您的 Go 代码。

## 构建

要构建可再分发的生产模式包，请使用 `wails build`。

## 下载

Linux：[hosts-switch-linux-amd64.deb](https://github.com/conkayyan/hosts-switch/releases/latest/download/hosts-switch-linux-amd64.deb)、[hosts-switch-linux-amd64](https://github.com/conkayyan/hosts-switch/releases/latest/download/hosts-switch-linux-amd64)

Windows：[hosts-switch-windows-amd64](https://github.com/conkayyan/hosts-switch/releases/latest/download/hosts-switch-windows-amd64)、[hosts-switch-amd64-installer.exe](https://github.com/conkayyan/hosts-switch/releases/latest/download/hosts-switch-amd64-installer.exe)

MacOs： [hosts-switch.pkg](https://github.com/conkayyan/hosts-switch/releases/latest/download/hosts-switch.pkg)，[hosts-switch.app.zip](https://github.com/conkayyan/hosts-switch/releases/latest/download/hosts-switch.app.zip)