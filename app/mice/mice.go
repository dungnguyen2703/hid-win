package mice

import "unsafe"

type BUTTON_ACTION string

const (
	LEFT_BUTTON_DOWN    BUTTON_ACTION = "LEFT_BUTTON_DOWN"
	LEFT_BUTTON_UP      BUTTON_ACTION = "LEFT_BUTTON_UP"
	RIGHT_BUTTON_DOWN   BUTTON_ACTION = "RIGHT_BUTTON_DOWN"
	RIGHT_BUTTON_UP     BUTTON_ACTION = "RIGHT_BUTTON_UP"
	MIDDLE_BUTTON_DOWN  BUTTON_ACTION = "MIDDLE_BUTTON_DOWN"
	MIDDLE_BUTTON_UP    BUTTON_ACTION = "MIDDLE_BUTTON_UP"
	WHEEL_SCROLL_UP     BUTTON_ACTION = "WHEEL_SCROLL_UP"
	WHEEL_SCROLL_DOWN   BUTTON_ACTION = "WHEEL_SCROLL_DOWN"
	BACK_BUTTON_DOWN    BUTTON_ACTION = "BACK_BUTTON_DOWN"
	BACK_BUTTON_UP      BUTTON_ACTION = "BACK_BUTTON_UP"
	FORWARD_BUTTON_DOWN BUTTON_ACTION = "FORWARD_BUTTON_DOWN"
	FORWARD_BUTTON_UP   BUTTON_ACTION = "FORWARD_BUTTON_UP"
)

type MSLLHOOKSTRUCT struct {
	Point     struct{ X, Y int32 }
	MouseData uint32
	Flags     uint32
	Time      uint32
	ExtraInfo uintptr
}

func Check(nCode int, wParam uintptr, lParam uintptr) (BUTTON_ACTION, bool) {
	if nCode >= 0 {
		data := (*MSLLHOOKSTRUCT)(unsafe.Pointer(lParam))
		switch wParam {
		case 0x0201:
			return LEFT_BUTTON_DOWN, true
		case 0x0202:
			return LEFT_BUTTON_UP, true
		case 0x0204:
			return RIGHT_BUTTON_DOWN, true
		case 0x0205:
			return RIGHT_BUTTON_UP, true
		case 0x0207:
			return MIDDLE_BUTTON_DOWN, true
		case 0x0208:
			return MIDDLE_BUTTON_UP, true
		case 0x020A:
			delta := int16(data.MouseData >> 16)
			if delta > 0 {
				return WHEEL_SCROLL_UP, true
			} else {
				return WHEEL_SCROLL_DOWN, true
			}
		case 0x020B:
			button := data.MouseData >> 16
			switch button {
			case 1:
				return BACK_BUTTON_DOWN, true
			case 2:
				return FORWARD_BUTTON_DOWN, true
			}
		case 0x020C:
			button := data.MouseData >> 16
			switch button {
			case 1:
				return BACK_BUTTON_UP, true
			case 2:
				return FORWARD_BUTTON_UP, true
			}
		}
	}
	return "", false
}
