package event

import "syscall"

type Action int

const (
	WindowLeft  Action = 0
	WindowRight Action = 1
)

const (
	KEYEVENTF_KEYDOWN = 0x0000
	KEYEVENTF_KEYUP   = 0x0002
)

const (
	KEY_CTRL  = 0x11
	KEY_WIN   = 0x5B
	KEY_LEFT  = 0x25
	KEY_RIGHT = 0x27
)

var (
	user32     = syscall.NewLazyDLL("user32.dll")
	keybdEvent = user32.NewProc("keybd_event")
)

func KeyDown(vkey byte) {
	keybdEvent.Call(uintptr(vkey), 0, KEYEVENTF_KEYDOWN, 0)
}

func KeyUp(vkey byte) {
	keybdEvent.Call(uintptr(vkey), 0, KEYEVENTF_KEYUP, 0)
}

func Run(action Action) {
	switch action {
	case WindowLeft:
		KeyDown(KEY_CTRL)
		KeyDown(KEY_WIN)
		KeyDown(KEY_LEFT)
		KeyUp(KEY_LEFT)
		KeyUp(KEY_WIN)
		KeyUp(KEY_CTRL)
	case WindowRight:
		KeyDown(KEY_CTRL)
		KeyDown(KEY_WIN)
		KeyDown(KEY_RIGHT)
		KeyUp(KEY_RIGHT)
		KeyUp(KEY_WIN)
		KeyUp(KEY_CTRL)
	}
}
