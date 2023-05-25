package main

import (
	"syscall"
	"unsafe"
)

var (
	user32       = syscall.NewLazyDLL("user32.dll")
	setCursorPos = user32.NewProc("SetCursorPos")
	getCursorPos = user32.NewProc("GetCursorPos")
	mouseEvent   = user32.NewProc("mouse_event")
)

const (
	MOUSEEVENTF_LEFTDOWN   = 0x0002
	MOUSEEVENTF_LEFTUP     = 0x0004
	MOUSEEVENTF_RIGHTDOWN  = 0x0008
	MOUSEEVENTF_RIGHTUP    = 0x0010
	MOUSEEVENTF_MIDDLEDOWN = 0x0020
	MOUSEEVENTF_MIDDLEUP   = 0x0040
)

type POINT struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func SetCursorPosition(x, y int) { //TODO: fix acceleration bug.
	currX, currY := GetCursorPosition()
	x = x + currX
	y = y + currY
	setCursorPos.Call(uintptr(x), uintptr(y))
}

func GetCursorPosition() (int, int) {
	var point POINT
	getCursorPos.Call(uintptr(unsafe.Pointer(&point)))
	return int(point.X), int(point.Y)
}

func LeftClick() {
	mouseEvent.Call(
		uintptr(MOUSEEVENTF_LEFTDOWN),
		0,
		0,
		0,
		0,
	)
	mouseEvent.Call(
		uintptr(MOUSEEVENTF_LEFTUP),
		0,
		0,
		0,
		0,
	)
}
