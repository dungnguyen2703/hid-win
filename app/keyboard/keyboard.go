package keyboard

import (
	"syscall"
	"unsafe"
)

const (
	WM_KEYDOWN    = 0x0100
	WM_SYSKEYDOWN = 0x0104
)

type KEY string

const (
	CTRL  KEY = "CTRL"
	ALT   KEY = "ALT"
	SHIFT KEY = "SHIFT"
	F1    KEY = "F1"
	F2    KEY = "F2"
	F3    KEY = "F3"
	F4    KEY = "F4"
	F5    KEY = "F5"
	F6    KEY = "F6"
	F7    KEY = "F7"
	F8    KEY = "F8"
	F9    KEY = "F9"
	F10   KEY = "F10"
	F11   KEY = "F11"
	F12   KEY = "F12"
)

type KBDLLHOOKSTRUCT struct {
	VkCode    uint32
	ScanCode  uint32
	Flags     uint32
	Time      uint32
	ExtraInfo uintptr
}

var (
	user32      = syscall.NewLazyDLL("user32.dll")
	getKeyState = user32.NewProc("GetKeyState")
)

func IsSystemKeyPressed(vKey int) bool {
	ret, _, _ := getKeyState.Call(uintptr(vKey))
	return ret&0x8000 != 0
}

func Check(nCode int, wParam uintptr, lParam uintptr) (KEY, []KEY, bool) {
	if nCode >= 0 && (wParam == WM_KEYDOWN || wParam == WM_SYSKEYDOWN) {
		kbData := (*KBDLLHOOKSTRUCT)(unsafe.Pointer(lParam))
		isCtrl := IsSystemKeyPressed(0x11)
		isAlt := IsSystemKeyPressed(0x12)
		isShift := IsSystemKeyPressed(0x10)
		var modifiers []KEY = []KEY{}
		if isCtrl {
			modifiers = append(modifiers, CTRL)
		}
		if isAlt {
			modifiers = append(modifiers, ALT)
		}
		if isShift {
			modifiers = append(modifiers, SHIFT)
		}
		switch kbData.VkCode {
		case 0x70:
			return F1, modifiers, true
		case 0x71:
			return F2, modifiers, true
		case 0x72:
			return F3, modifiers, true
		case 0x73:
			return F4, modifiers, true
		case 0x74:
			return F5, modifiers, true
		case 0x75:
			return F6, modifiers, true
		case 0x76:
			return F7, modifiers, true
		case 0x77:
			return F8, modifiers, true
		case 0x78:
			return F9, modifiers, true
		case 0x79:
			return F10, modifiers, true
		case 0x7A:
			return F11, modifiers, true
		case 0x7B:
			return F12, modifiers, true
		}
		return "", modifiers, true
	}
	return "", nil, false
}
