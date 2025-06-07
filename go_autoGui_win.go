//go:build windows

package goAutoGui

import (
	"log"
	"syscall"
	"time"
	"unsafe"
)

// SystemLang 系统语言类型 zh,en
var SystemLang string

var user32DLL = syscall.MustLoadDLL(`User32.dll`)
var mouseEventProc = user32DLL.MustFindProc("mouse_event")

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
		return nil, err
	}
	var pt POINT
	r1, _, err := proc.Call(uintptr(unsafe.Pointer(&pt)))
	if r1 == 0 {
		return nil, err
	}
	return &pt, nil
}

// var MouseEventMove = 0x0001       // 移动鼠标           (十):1

var MouseEventABSOLUTE = 0x8000 // 标示是否采取绝对坐标    (十):32768

// MoveMouse 平滑移动鼠标,speed毫秒到达,每次移动间隔step毫秒
func MoveMouse(x, y int32, speed, step int) error {
	pos, err := GetCursorPos()
	if err != nil {
		return err
	}
	cursor := user32DLL.MustFindProc("SetCursorPos")
	stepDuration := time.Millisecond * time.Duration(step)
	totalSteps := int(float64(speed) / float64(stepDuration.Milliseconds()))
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
	var mouseEventLeftDown = 0x0002  // 模仿鼠标左键按下    (十):2
	var mouseEventRightDown = 0x0008 // 模仿鼠标右键按下    (十):8

	var event int
	if left {
		event = mouseEventLeftDown
	} else {
		event = mouseEventRightDown
	}
	ret, _, err := mouseEventProc.Call(
		uintptr(event), // 事件
		0,              // x相对位置
		0,              // y相对位置
		0,              // 滚轮
		0,
	)
	if ret == 0 {
		return err
	}

	return nil
}

// MouseUp 鼠标{松开} left是true的时候左键，反之右键
func MouseUp(left bool) error {
	var mouseEventLeftUp = 0x0004  // 模仿鼠标左键抬起    (十):4
	var mouseEventRightUp = 0x0010 // 模仿鼠标右键抬起    (十):16
	var event int
	if left {
		event = mouseEventLeftUp
	} else {
		event = mouseEventRightUp
	}
	ret, _, err := mouseEventProc.Call(
		uintptr(event), // 事件
		0,              // x相对位置
		0,              // y相对位置
		0,              // 滚轮
		0,
	)
	if ret == 0 {
		return err
	}
	return nil
}

// MouseWheelDown 鼠标中建按下 左倾斜(<0) 中键(=0) 右倾斜(>0)
func MouseWheelDown() error {
	var MouseEventMiddleDown = 0x0020 // 模仿鼠标中键按下	(十):32
	ret, _, err := mouseEventProc.Call(
		uintptr(MouseEventMiddleDown), // 事件
		0,                             // x相对位置
		0,                             // y相对位置
		0,                             // 滚轮
		0,
	)
	if ret == 0 {
		return err
	}
	return nil
}

// MouseWheelUp 鼠标中键松开 左倾斜(<0) 中键(=0) 右倾斜(>0)
func MouseWheelUp(left int) error {
	var MouseEventMiddleUp = 0x0040 // 模仿鼠标中键松开	(十):64
	ret, _, err := mouseEventProc.Call(
		uintptr(MouseEventMiddleUp), // 事件
		0,                           // x相对位置
		0,                           // y相对位置
		0,                           // 滚轮
		0,
	)
	if ret == 0 {
		return err
	}
	return nil
}

// MouseMiddleWheel 滚轮向上
func MouseMiddleWheel(Upward bool) error {
	return nil
}

// Click 点击某处
func Click(times ClickTimes) error {
	if times.times <= 0 {
		times.times = 1
	}
	for i := uint(0); i < times.times; i++ {
		err := MouseDown(true)
		if err != nil {
			log.Fatal(err)
		}
		time.Sleep(time.Millisecond * 100)
		err = MouseUp(true)
		if err != nil {
			log.Fatal(err)
		}
	}
	return nil
}

func MoveAndClick(x, y int32, other ClickStruct) error {
	err := MoveMouse(x, y, other.Speed, other.step)
	if err != nil {
		return err
	}
	err = Click(other.times)
	if err != nil {
		return err
	}
	return nil
}
