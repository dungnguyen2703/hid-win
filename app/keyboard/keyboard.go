package keyboard

import (
	"hidtool/app/logger"
	"unsafe"
)

const (
	WM_KEYDOWN              = 0x0100
	WM_KEYUP                = 0x0101
	WM_SYSKEYDOWN           = 0x0104
	WM_SYSKEYUP             = 0x0105
	LLKHF_INJECTED          = 0x00000010
	LLKHF_LOWER_IL_INJECTED = 0x00000002
	LLKHF_ALTDOWN           = 0x00000020
	LLKHF_EXTENDED          = 0x00000001
)

type KBDLLHOOKSTRUCT struct {
	VkCode    uint32
	ScanCode  uint32
	Flags     uint32
	Time      uint32
	ExtraInfo uintptr
}

var keyStates [256]bool

// Check keyboard event and return key, eventType, isFirst, isExtended, isAltDown, ok
// Key: the key that triggered the event
// eventType: TAP_DOWN or TAP_UP
// isFirst: whether the key event is the first TAP_DOWN (not repeated due to key hold)
// isExtended: whether the key is an extended key
// isAltDown: whether the Alt key was held down during the event
// ok: whether event is valid and the hardware key was tapped (not injected)
func Check(nCode int, wParam uintptr, lParam uintptr) (KEY, ACTION, bool, bool, bool, bool) {
	if nCode < 0 {
		return 0, "", false, false, false, false
	}
	var eventType ACTION
	switch wParam {
	case WM_KEYDOWN, WM_SYSKEYDOWN:
		eventType = TAP_DOWN
	case WM_KEYUP, WM_SYSKEYUP:
		eventType = TAP_UP
	default:
		return 0, "", false, false, false, false
	}
	kbData := (*KBDLLHOOKSTRUCT)(unsafe.Pointer(lParam))
	vkCode := kbData.VkCode
	if vkCode > 0xFF {
		return 0, "", false, false, false, false
	}
	isFirst := true
	switch eventType {
	case TAP_DOWN:
		if keyStates[vkCode] {
			isFirst = false
		}
		keyStates[vkCode] = true
	case TAP_UP:
		keyStates[vkCode] = false
	}
	isInjected := (kbData.Flags&LLKHF_INJECTED) != 0 || (kbData.Flags&LLKHF_LOWER_IL_INJECTED) != 0
	if isInjected {
		return 0, "", isFirst, false, false, false
	}
	isExtended := (kbData.Flags & LLKHF_EXTENDED) != 0
	isAltDown := (kbData.Flags & LLKHF_ALTDOWN) != 0
	key, ok := getKey(vkCode)
	if !ok {
		logger.Debug(key, "Pressed")
	}
	return key, eventType, isFirst, isExtended, isAltDown, ok
}

func Reset() {
	keyStates = [256]bool{}
}

func IsKeyPressed(vKey KEY) bool {
	return keyStates[vKey]
}
