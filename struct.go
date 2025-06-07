package goAutoGui

// ScreenSizeStruct 定义了屏幕大小的类
type ScreenSizeStruct struct {
	Width  uintptr // 屏幕宽度
	Height uintptr // 屏幕高度
}

type POINT struct {
	X int32
	Y int32
}
type ClickTimes struct {
	times uint
}
type ClickStruct struct {
	Times       uint
	times       ClickTimes
	Speed, step int
}
