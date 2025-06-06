<p align="center">
  <a href="#">
    <img src="assets/icon.png" width="200" height="200" alt="go-autoGUI">
  </a>
</p>

<div align="center">

# go-autoGUI

_✨ Golang 原生实现 write by shaoze 少泽 ✨_

</div>

## 🌟 项目介绍 🌟

有些老程序不支持控件操控，但是目前(2025.06)没有看见有基于图像识别的自动化程序，可能以后有但是目前没有，所以项目启动这个项目
。

虽然 [robotGo](https://pkg.go.dev/github.com/go-vgo/robotgo) 确实很强大,但是类似[PyAutoGUI](https://pypi.org/project/PyAutoGUI/) 机制的内容几乎没有

(于是狠狠的往里面注入了福瑞代码)

## 兼容性


|   系统    | 支持 |
|:-------:|:--:|
| windows | ✅  |
|  other  | ❌  |


> ⚠️ go_autoGUI 目前仅支持 windows (有考虑后期加入其他系统支持)

## 开始

安装
```
go install xxxxx
```
示例：
```
package main


```

## 文档

### 结构体

### 函数

#### ScreenSize 获取屏幕大小

示例：

|   字段    | 类型 |
|:-------:|:--:|
| :return | struct  |

```golang
size, err := go_autoGUI.ScreenSize()
if err != nil {
	return
}
fmt.Println(size.Width, size.Height)
```
