package w

import (
	"log"
	"syscall"
	"unsafe"
)

type ActionType int

const (
	WindowLeft  ActionType = 0
	WindowRight ActionType = 1
)

var (
	user32           = syscall.NewLazyDLL("user32.dll")
	setWindowsHookEx = user32.NewProc("SetWindowsHookExW")
	callNextHookEx   = user32.NewProc("CallNextHookEx")
	keybdEvent       = user32.NewProc("keybd_event")
	getMessage       = user32.NewProc("GetMessageW")
)

const (
	WH_MOUSE_LL        = 14
	LEFT_BUTTON_DOWN   = 0x0201
	LEFT_BUTTON_UP     = 0x0202
	RIGHT_BUTTON_DOWN  = 0x0204
	RIGHT_BUTTON_UP    = 0x0205
	MIDDLE_BUTTON_DOWN = 0x0207
	MIDDLE_BUTTON_UP   = 0x0208
	WHEEL_SCROLL       = 0x020A
	XBUTTON_DOWN       = 0x020B // 4: Back, 5: Forward
	XBUTTON_UP         = 0x020C // 4: Back, 5: Forward
	KEYEVENTF_KEYDOWN  = 0x0000
	KEYEVENTF_KEYUP    = 0x0002
)

const (
	KEY_CTRL  = 0x11
	KEY_WIN   = 0x5B
	KEY_LEFT  = 0x25
	KEY_RIGHT = 0x27
)

type MSLLHOOKSTRUCT struct {
	Point     struct{ X, Y int32 }
	MouseData uint32
	Flags     uint32
	Time      uint32
	ExtraInfo uintptr
}

func debugLog(v ...any) {
	log.Println(v...)
}

func Run() {
	hook, _, _ := setWindowsHookEx.Call(
		WH_MOUSE_LL,
		syscall.NewCallback(func(nCode int, wParam uintptr, lParam uintptr) uintptr {
			if nCode >= 0 {
				data := (*MSLLHOOKSTRUCT)(unsafe.Pointer(lParam))
				switch wParam {
				case LEFT_BUTTON_DOWN:
					debugLog("🖱️ Left Button Down")
				case LEFT_BUTTON_UP:
					debugLog("🖱️ Left Button Up")
				case RIGHT_BUTTON_DOWN:
					debugLog("🖱️ Right Button Down")
				case RIGHT_BUTTON_UP:
					debugLog("🖱️ Right Button Up")
				case MIDDLE_BUTTON_DOWN:
					debugLog("🖱️ Middle Button Down")
				case MIDDLE_BUTTON_UP:
					debugLog("🖱️ MiddleButton  Up")
				case WHEEL_SCROLL:
					delta := int16(data.MouseData >> 16)
					if delta > 0 {
						debugLog("🖱️ Wheel Scroll Up")
					} else {
						debugLog("🖱️ Wheel Scroll Down")
					}
				case XBUTTON_DOWN:
					button := data.MouseData >> 16
					switch button {
					case 1:
						debugLog("🔙 Back Button Down")
						doAction(WindowLeft)
						return 1 // Cancel back event
					case 2:
						debugLog("🔜 Forward Button Down")
						doAction(WindowRight)
						return 1 // Cancel forward event
					}
				case XBUTTON_UP:
					button := data.MouseData >> 16
					switch button {
					case 1:
						debugLog("🔙 Back Button Up")
						return 1 // Cancel back event
					case 2:
						debugLog("🔜 Forward Button Up")
						return 1 // Cancel forward event
					}
				}
			}
			ret, _, _ := callNextHookEx.Call(0, uintptr(nCode), wParam, lParam)
			return ret
		}),
		0,
		0,
	)

	// Loop to keep the hook alive
	var msg struct {
		Context uintptr
		Msg     uint32
		WParam  uintptr
		LParam  uintptr
		Time    uint32
		Pt      struct{ X, Y int32 }
	}
	for {
		ret, _, _ := getMessage.Call(uintptr(unsafe.Pointer(&msg)), 0, 0, 0)
		if ret == 0 {
			break
		}
	}
	_ = hook
}

func keyDown(vkey byte) {
	debugLog("🔘 Key Down:", vkey)
	keybdEvent.Call(uintptr(vkey), 0, KEYEVENTF_KEYDOWN, 0)
}

func keyUp(vkey byte) {
	debugLog("🔘 Key Up:", vkey)
	keybdEvent.Call(uintptr(vkey), 0, KEYEVENTF_KEYUP, 0)
}

func keyPress(vkey byte) {
	debugLog("🔘 Key Press:", vkey)
	keybdEvent.Call(uintptr(vkey), 0, 0, 0)
	keybdEvent.Call(uintptr(vkey), 0, KEYEVENTF_KEYUP, 0)
}

func doAction(action ActionType) {
	switch action {
	case WindowLeft:
		keyDown(KEY_CTRL)
		keyDown(KEY_WIN)
		keyDown(KEY_LEFT)
		keyUp(KEY_LEFT)
		keyUp(KEY_WIN)
		keyUp(KEY_CTRL)
	case WindowRight:
		keyDown(KEY_CTRL)
		keyDown(KEY_WIN)
		keyDown(KEY_RIGHT)
		keyUp(KEY_RIGHT)
		keyUp(KEY_WIN)
		keyUp(KEY_CTRL)
	}
}
