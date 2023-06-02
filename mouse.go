package main

import (
	"fmt"
	"log"
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
	X int32 `json:"x"`
	Y int32 `json:"y"`
}

func SetCursorPosition(x, y int32) {
	currX, currY, _ := GetCursorPosition()
	x = x/10 + currX
	y = y/10 + currY

	setCursorPos.Call(uintptr(x), uintptr(y))
}

func GetCursorPosition() (int32, int32, error) {
	var point POINT
	ret, _, _ := getCursorPos.Call(uintptr(unsafe.Pointer(&point)))
	if ret == 0 {
		log.Printf("failed to get cursor position")
		return 0, 0, fmt.Errorf("failed to get cursor position")
	}

	return int32(point.X), int32(point.Y), nil
}

func LeftClick() {
	mouseEvent.Call(
		uintptr(MOUSEEVENTF_LEFTDOWN),
	)

	mouseEvent.Call(
		uintptr(MOUSEEVENTF_LEFTUP),
	)
}

func RightClick() {
	mouseEvent.Call(
		uintptr(MOUSEEVENTF_RIGHTDOWN),
	)

	mouseEvent.Call(
		uintptr(MOUSEEVENTF_RIGHTUP),
	)
}
