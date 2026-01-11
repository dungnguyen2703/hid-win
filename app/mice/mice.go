package mice

import (
	"hidtool/app/logger"
	"unsafe"
)

const (
	WM_MOUSEMOVE   = 0x0200
	WM_LBUTTONDOWN = 0x0201
	WM_LBUTTONUP   = 0x0202
	WM_RBUTTONDOWN = 0x0204
	WM_RBUTTONUP   = 0x0205
	WM_MBUTTONDOWN = 0x0207
	WM_MBUTTONUP   = 0x0208
	WM_MOUSEWHEEL  = 0x020A
	WM_XBUTTONDOWN = 0x020B
	WM_XBUTTONUP   = 0x020C
	WM_MOUSEHWHEEL = 0x020E

	LLMHF_INJECTED          = 0x00000001
	LLMHF_LOWER_IL_INJECTED = 0x00000002
)

type MSLLHOOKSTRUCT struct {
	Point     struct{ X, Y int32 }
	MouseData uint32
	Flags     uint32
	Time      uint32
	ExtraInfo uintptr
}

var buttonStates = make(map[BUTTON]bool)

func Check(nCode int, wParam uintptr, lParam uintptr) (BUTTON, ACTION, bool, bool) {
	if nCode < 0 {
		return "", "", false, false
	}
	data := (*MSLLHOOKSTRUCT)(unsafe.Pointer(lParam))
	isInjected := (data.Flags&LLMHF_INJECTED) != 0 || (data.Flags&LLMHF_LOWER_IL_INJECTED) != 0

	switch wParam {
	case WM_LBUTTONDOWN:
		buttonStates[LEFT_BUTTON] = true
		return LEFT_BUTTON, CLICK_DOWN, true, isInjected
	case WM_LBUTTONUP:
		buttonStates[LEFT_BUTTON] = false
		return LEFT_BUTTON, CLICK_UP, true, isInjected
	case WM_RBUTTONDOWN:
		buttonStates[RIGHT_BUTTON] = true
		return RIGHT_BUTTON, CLICK_DOWN, true, isInjected
	case WM_RBUTTONUP:
		buttonStates[RIGHT_BUTTON] = false
		return RIGHT_BUTTON, CLICK_UP, true, isInjected
	case WM_MBUTTONDOWN:
		buttonStates[MIDDLE_BUTTON] = true
		return MIDDLE_BUTTON, CLICK_DOWN, true, isInjected
	case WM_MBUTTONUP:
		buttonStates[MIDDLE_BUTTON] = false
		return MIDDLE_BUTTON, CLICK_UP, true, isInjected
	case WM_MOUSEWHEEL:
		delta := int16(data.MouseData >> 16)
		if delta > 0 {
			return V_WHEEL, SCROLL_UP, true, isInjected
		} else {
			return V_WHEEL, SCROLL_DOWN, true, isInjected
		}
	case WM_MOUSEHWHEEL:
		delta := int16(data.MouseData >> 16)
		if delta > 0 {
			return H_WHEEL, SCROLL_RIGHT, true, isInjected
		} else {
			return H_WHEEL, SCROLL_LEFT, true, isInjected
		}
	case WM_XBUTTONDOWN:
		button := data.MouseData >> 16
		switch button {
		case 1:
			buttonStates[BACK_BUTTON] = true
			return BACK_BUTTON, CLICK_DOWN, true, isInjected
		case 2:
			buttonStates[FORWARD_BUTTON] = true
			return FORWARD_BUTTON, CLICK_DOWN, true, isInjected
		}
	case WM_XBUTTONUP:
		button := data.MouseData >> 16
		switch button {
		case 1:
			buttonStates[BACK_BUTTON] = false
			return BACK_BUTTON, CLICK_UP, true, isInjected
		case 2:
			buttonStates[FORWARD_BUTTON] = false
			return FORWARD_BUTTON, CLICK_UP, true, isInjected
		}
	}
	if wParam == WM_MOUSEMOVE {
		// Mouse Move event
	} else {
		logger.Error("unknown mouse event:", nCode, wParam)
	}
	return "", "", false, isInjected
}

func Reset() {
	buttonStates = make(map[BUTTON]bool)
}

func IsButtonClicked(btn BUTTON) bool {
	state, ok := buttonStates[btn]
	if !ok {
		return false
	}
	return state
}
