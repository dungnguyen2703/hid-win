package keyboard

import "fmt"

type KEY uint

const (
	A KEY = 0x41
	B KEY = 0x42
	C KEY = 0x43
	D KEY = 0x44
	E KEY = 0x45
	F KEY = 0x46
	G KEY = 0x47
	H KEY = 0x48
	I KEY = 0x49
	J KEY = 0x4A
	K KEY = 0x4B
	L KEY = 0x4C
	M KEY = 0x4D
	N KEY = 0x4E
	O KEY = 0x4F
	P KEY = 0x50
	Q KEY = 0x51
	R KEY = 0x52
	S KEY = 0x53
	T KEY = 0x54
	U KEY = 0x55
	V KEY = 0x56
	W KEY = 0x57
	X KEY = 0x58
	Y KEY = 0x59
	Z KEY = 0x5A
)

const (
	N_0     KEY = 0x30
	N_1     KEY = 0x31
	N_2     KEY = 0x32
	N_3     KEY = 0x33
	N_4     KEY = 0x34
	N_5     KEY = 0x35
	N_6     KEY = 0x36
	N_7     KEY = 0x37
	N_8     KEY = 0x38
	N_9     KEY = 0x39
	NUM_0   KEY = 0x60
	NUM_1   KEY = 0x61
	NUM_2   KEY = 0x62
	NUM_3   KEY = 0x63
	NUM_4   KEY = 0x64
	NUM_5   KEY = 0x65
	NUM_6   KEY = 0x66
	NUM_7   KEY = 0x67
	NUM_8   KEY = 0x68
	NUM_9   KEY = 0x69
	NUM_MUL KEY = 0x6A // *
	NUM_ADD KEY = 0x6B // +
	NUM_SUB KEY = 0x6D // -
	NUM_DEC KEY = 0x6E // .
	NUM_DIV KEY = 0x6F // /
)

const (
	F1  KEY = 0x70
	F2  KEY = 0x71
	F3  KEY = 0x72
	F4  KEY = 0x73
	F5  KEY = 0x74
	F6  KEY = 0x75
	F7  KEY = 0x76
	F8  KEY = 0x77
	F9  KEY = 0x78
	F10 KEY = 0x79
	F11 KEY = 0x7A
	F12 KEY = 0x7B
)

const (
	BACKSPACE    KEY = 0x08
	TAB          KEY = 0x09
	ENTER        KEY = 0x0D
	SHIFT        KEY = 0x10
	CTRL         KEY = 0x11
	ALT          KEY = 0x12
	PAUSE        KEY = 0x13
	CAPS_LOCK    KEY = 0x14
	ESCAPE       KEY = 0x1B
	SPACE        KEY = 0x20
	PAGE_UP      KEY = 0x21
	PAGE_DOWN    KEY = 0x22
	END          KEY = 0x23
	HOME         KEY = 0x24
	ARROW_LEFT   KEY = 0x25
	ARROW_UP     KEY = 0x26
	ARROW_RIGHT  KEY = 0x27
	ARROW_DOWN   KEY = 0x28
	PRINT_SCREEN KEY = 0x2C
	INSERT       KEY = 0x2D
	DELETE       KEY = 0x2E
	WIN          KEY = 0x5B
	RIGHT_WIN    KEY = 0x5C
	NUM_LOCK     KEY = 0x90
	SCROLL_LOCK  KEY = 0x91
	LEFT_SHIFT   KEY = 0xA0
	RIGHT_SHIFT  KEY = 0xA1
	LEFT_CTRL    KEY = 0xA2
	RIGHT_CTRL   KEY = 0xA3
	LEFT_ALT     KEY = 0xA4
	RIGHT_ALT    KEY = 0xA5
)

const (
	BRIGHTNESS_DOWN KEY = 0x81
	BRIGHTNESS_UP   KEY = 0x82
)

const (
	BROWSER_BACK      KEY = 0xA6
	BROWSER_FORWARD   KEY = 0xA7
	BROWSER_REFRESH   KEY = 0xA8
	BROWSER_STOP      KEY = 0xA9
	BROWSER_SEARCH    KEY = 0xAA
	BROWSER_FAVORITES KEY = 0xAB
	BROWSER_HOME      KEY = 0xAC
)

const (
	LAUNCH_MAIL         KEY = 0xB4
	LAUNCH_MEDIA_SELECT KEY = 0xB5
	LAUNCH_APP1         KEY = 0xB6
	LAUNCH_APP2         KEY = 0xB7
)

const (
	OEM_1      KEY = 0xBA // ; :
	OEM_PLUS   KEY = 0xBB // = +
	OEM_COMMA  KEY = 0xBC // , <
	OEM_MINUS  KEY = 0xBD // - _
	OEM_PERIOD KEY = 0xBE // . >
	OEM_2      KEY = 0xBF // / ?
	OEM_3      KEY = 0xC0 // ` ~
	OEM_4      KEY = 0xDB // [ {
	OEM_5      KEY = 0xDC // \ |
	OEM_6      KEY = 0xDD // ] }
	OEM_7      KEY = 0xDE // ' "
)

const (
	VOLUME_MUTE KEY = 0xAD
	VOLUME_DOWN KEY = 0xAE
	VOLUME_UP   KEY = 0xAF
	MEDIA_NEXT  KEY = 0xB0
	MEDIA_PREV  KEY = 0xB1
	MEDIA_STOP  KEY = 0xB2
	MEDIA_PLAY  KEY = 0xB3
)

const (
	CONTEXT_MENU KEY = 0x5D
	SLEEP        KEY = 0x5F
	SELECT       KEY = 0x29
	EXECUTE      KEY = 0x2B
	HELP         KEY = 0x2F
)

func getKey(vkCode uint32) (KEY, bool) {
	key := KEY(vkCode)
	_, exists := keyMap[key]
	return key, exists
}

func (m KEY) String() string {
	if name, exists := keyMap[m]; exists {
		return name
	}
	return fmt.Sprintf("KEY (0x%X)", uint32(m))
}

type ACTION string

const (
	TAP_DOWN ACTION = "TAP_DOWN"
	TAP_UP   ACTION = "TAP_UP"
)
