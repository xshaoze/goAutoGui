//go:build windows

package goAutoGUI

import (
	"syscall"
	"time"
	"unsafe"
)

// SystemLang 系统语言类型 zh,en
var SystemLang string

var user32DLL = syscall.MustLoadDLL(`User32.dll`)

func init() {
	defer func(user32DLL *syscall.DLL) {
		err := user32DLL.Release()
		if err != nil {
		}
	}(user32DLL)
}

// ScreenSize 获取屏幕大小
func ScreenSize() (ScreenSizeStruct, error) {
	getSystemMetricsProc := user32DLL.MustFindProc(`GetSystemMetrics`)
	width, _, _ := getSystemMetricsProc.Call(uintptr(0))
	height, _, _ := getSystemMetricsProc.Call(uintptr(1))
	return ScreenSizeStruct{
		width,
		height,
	}, nil
}

// GetCursorPos 获取鼠标当前位置
func GetCursorPos() (*POINT, error) {
	// 安全地获取 API 函数指针
	proc, err := user32DLL.FindProc("GetCursorPos")
	if err != nil {
	}

	var pt POINT
	r1, _, err := proc.Call(uintptr(unsafe.Pointer(&pt)))

	if r1 == 0 {
	}

	return &pt, nil
}

// MoveMouse 平滑移动鼠标,speed毫秒到达,每次移动间隔step毫秒
func MoveMouse(x, y int32, speed, step int) error {
	pos, err := GetCursorPos()
	if err != nil {
		return err
	}
	cursor := user32DLL.MustFindProc("SetCursorPos")
	// 计算每一步的时间间隔（1 毫秒）
	stepDuration := time.Millisecond * time.Duration(step)
	// 计算总步数
	totalSteps := int(float64(speed) / float64(stepDuration.Milliseconds()))
	// 计算每一步的增量
	dx := float64(x-pos.X) / float64(totalSteps)
	dy := float64(y-pos.Y) / float64(totalSteps)

	currentX, currentY := float64(pos.X), float64(pos.Y)

	for i := 0; i <= totalSteps; i++ {
		_, _, _ = cursor.Call(uintptr(currentX), uintptr(currentY))
		time.Sleep(stepDuration)
		currentX += dx
		currentY += dy
	}
	return nil
}

// MouseDown 鼠标{按下} left是true的时候左键，反之右键
func MouseDown(left bool) error {
	return nil
}

// MouseUp 鼠标{松开} left是true的时候左键，反之右键
func MouseUp(left bool) error {
	return nil
}

// MouseWheelDown 鼠标中建按下 左倾斜(<0) 中键(=0) 右倾斜(>0)
func MouseWheelDown(left int) error {
	return nil
}

// MouseWheelUp 鼠标中键松开 左倾斜(<0) 中键(=0) 右倾斜(>0)
func MouseWheelUp(left int) error {
	return nil
}

// MouseOtherDown 自定义鼠标键{按下}
func MouseOtherDown() error {
	return nil
}

// MouseOtherUp 自定义鼠标键{松开}
func MouseOtherUp() error {
	return nil
}

// MouseMiddleWheel 滚轮向上
func MouseMiddleWheel(Upward bool) error {
	return nil
}

// Click 点击某处
func Click(times uint) error {
	return nil
}

func MoveAndClick(c ClickStruct) error {
	// 默认值
	if c.times == 0 {
		c.times = 1
	}
	if c.step == 0 {
		c.step = 10
	}
	if c.speed == 0 {
		c.speed = 100
	}

	err := MoveMouse(c.x, c.y, c.speed, c.step)
	if err != nil {
		return err
	}
	err = Click(c.times)
	if err != nil {
		return err
	}
	return nil
}
