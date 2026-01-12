package app

import (
	"hidtool/app/keyboard"
	"hidtool/app/logger"
	"hidtool/app/mice"
	"hidtool/app/profile"
	"syscall"
	"time"
	"unsafe"
)

var (
	user32           = syscall.NewLazyDLL("user32.dll")
	setWindowsHookEx = user32.NewProc("SetWindowsHookExW")
	callNextHookEx   = user32.NewProc("CallNextHookEx")
	getMessage       = user32.NewProc("GetMessageW")
	openInputDesktop = user32.NewProc("OpenInputDesktop")
	closeDesktop     = user32.NewProc("CloseDesktop")
)

const (
	WH_KEYBOARD_LL = 13
	WH_MOUSE_LL    = 14
)

func Run() {
	hook, _, _ := setWindowsHookEx.Call(
		WH_MOUSE_LL,
		syscall.NewCallback(func(nCode int, wParam uintptr, lParam uintptr) uintptr {
			currentProfile := profile.GetCurrentProfile()
			if currentProfile != nil {
				key, eventType, ok, isInjected := mice.Check(nCode, wParam, lParam)
				if ok && eventType == mice.CLICK_DOWN && !isInjected {
					logger.Debug("🐭 Mouse Event:", key)
					binding := currentProfile.GetBinding(0, key)
					if binding != nil {
						binding.Action()
						if binding.DisableLatestInput() {
							return 1
						}
					}
				}
			}
			ret, _, _ := callNextHookEx.Call(0, uintptr(nCode), wParam, lParam)
			return ret
		}),
		0,
		0,
	)
	kbHook, _, _ := setWindowsHookEx.Call(
		WH_KEYBOARD_LL,
		syscall.NewCallback(func(nCode int, wParam uintptr, lParam uintptr) uintptr {
			currentProfile := profile.GetCurrentProfile()
			if currentProfile != nil {
				key, eventType, isFirst, isExtended, _, ok := keyboard.Check(nCode, wParam, lParam)
				if ok && eventType == keyboard.TAP_DOWN {
					if isExtended {
						logger.Debug("⌨️ Extended Keyboard Event:", key)
					} else {
						logger.Debug("⌨️ Keyboard Event:", key)
					}

					binding := currentProfile.GetBinding(key, "")
					if binding != nil {
						if isFirst {
							binding.Action()
						}
						if binding.DisableLatestInput() {
							return 1
						}
					}
				}
			}
			ret, _, _ := callNextHookEx.Call(0, uintptr(nCode), wParam, lParam)
			return ret
		}),
		0,
		0,
	)

	go func() {
		wasLocked := false

		for {
			time.Sleep(2 * time.Second)
			hDesk, _, _ := openInputDesktop.Call(0, 0, 0x100)
			isLocked := (hDesk == 0)
			if hDesk != 0 {
				closeDesktop.Call(hDesk)
			}
			if isLocked {
				wasLocked = true
			} else if wasLocked {
				keyboard.Reset()
				mice.Reset()
				wasLocked = false
			}
		}
	}()

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
	_ = kbHook
}
