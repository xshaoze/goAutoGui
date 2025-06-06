//go:build windows

package go_autoGUI

import (
	"goAutoGui/internal/appError"
	"os/exec"
	"strings"
	"syscall"
	"unsafe"
)

// SystemLang 系统语言类型 zh,en
var SystemLang string

var User32DLL = syscall.MustLoadDLL(`User32.dll`)

func init() {
	// 获取语言
	cmd := exec.Command("powershell", "Get-Culture | select -exp Name")
	output, err := cmd.Output()
	if err == nil {
		if lang := strings.TrimSpace(string(output)); lang == "zh-CN" {
			SystemLang = "zh"
		} else {
			SystemLang = "en"
		}
	}
	defer func(user32DLL *syscall.DLL) {
		err := user32DLL.Release()
		if err != nil {
		}
	}(User32DLL)
}

// ScreenSize 获取屏幕大小
func ScreenSize() (ScreenSizeStruct, error) {
	getSystemMetricsProc := User32DLL.MustFindProc(`GetSystemMetrics`)
	width, _, _ := getSystemMetricsProc.Call(uintptr(0))
	height, _, _ := getSystemMetricsProc.Call(uintptr(1))
	return ScreenSizeStruct{
		width,
		height,
	}, nil
}

// GetCursorPos 获取鼠标当前位置
func GetCursorPos() (POINT, error) {
	getCursorPos := User32DLL.MustFindProc("GetCursorPos")
	var pt POINT
	ret, _, _ := getCursorPos.Call(uintptr(unsafe.Pointer(&pt)))
	if ret == 0 {
		return POINT{}, appError.Error("getCursorPosError", SystemLang)
	}
	return pt, nil
}

// MoveMouse 平滑移动鼠标,speed秒到达
func MoveMouse(x, y int, speed float64) error {
	//cursor := User32DLL.MustFindProc("SetCursorPos")
	//
	//pos, err := GetCursorPos()
	//if err != nil {
	//	return err
	//}

	//pos.X, pos.Y
	//pyX := x - pos.X
	//pyY := y - pos.Y
	//k := pyX / pyY
	//b := y - k*x
	//sep := pyX / int(speed*10)
	//
	//for i := 0; i < int(speed*10); i++ {
	//	_, _, _ = cursor.Call(uintptr(sep*i), uintptr(k*sep*i+b))
	//	time.Sleep(100 * time.Millisecond)
	//}

	return nil
}

// Click 点击某处
func Click(x, y, times int, speed float64) {

}
