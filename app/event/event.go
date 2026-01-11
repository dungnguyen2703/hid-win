package event

import (
	"hidtool/app/keyboard"
	"syscall"
)

const (
	KEYEVENTF_KEYDOWN = 0x0000
	KEYEVENTF_KEYUP   = 0x0002
)

var (
	user32     = syscall.NewLazyDLL("user32.dll")
	keybdEvent = user32.NewProc("keybd_event")
)

func KeyDown(vkey keyboard.KEY) {
	keybdEvent.Call(uintptr(vkey), 0, KEYEVENTF_KEYDOWN, 0)
}

func KeyUp(vkey keyboard.KEY) {
	keybdEvent.Call(uintptr(vkey), 0, KEYEVENTF_KEYUP, 0)
}

func KeyPressed(vkey keyboard.KEY) {
	keybdEvent.Call(uintptr(vkey), 0, KEYEVENTF_KEYDOWN, 0)
	keybdEvent.Call(uintptr(vkey), 0, KEYEVENTF_KEYUP, 0)
}
