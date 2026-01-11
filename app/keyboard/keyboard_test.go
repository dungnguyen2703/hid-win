package keyboard

import (
	"fmt"
	"testing"
)

func TestKeyMapCompleteness(t *testing.T) {
	keysToCheck := []KEY{
		A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T, U, V, W, X, Y, Z,
		N_0, N_1, N_2, N_3, N_4, N_5, N_6, N_7, N_8, N_9,
		NUM_0, NUM_1, NUM_2, NUM_3, NUM_4, NUM_5, NUM_6, NUM_7, NUM_8, NUM_9,
		NUM_MUL, NUM_ADD, NUM_SUB, NUM_DEC, NUM_DIV,
		F1, F2, F3, F4, F5, F6, F7, F8, F9, F10, F11, F12,
		BACKSPACE, TAB, ENTER, SHIFT, CTRL, ALT, ESCAPE, SPACE,
		LEFT_SHIFT, RIGHT_SHIFT, LEFT_CTRL, RIGHT_CTRL, LEFT_ALT, RIGHT_ALT,
		OEM_1, OEM_2, OEM_3, OEM_4, OEM_5, OEM_6, OEM_7,
		OEM_PLUS, OEM_COMMA, OEM_MINUS, OEM_PERIOD,
		VOLUME_UP, VOLUME_DOWN, VOLUME_MUTE,
		MEDIA_NEXT, MEDIA_PREV, MEDIA_STOP, MEDIA_PLAY,
		BRIGHTNESS_DOWN, BRIGHTNESS_UP,
		BROWSER_BACK, BROWSER_FORWARD, BROWSER_REFRESH, BROWSER_STOP, BROWSER_SEARCH, BROWSER_FAVORITES, BROWSER_HOME,
		LAUNCH_MAIL, LAUNCH_MEDIA_SELECT, LAUNCH_APP1, LAUNCH_APP2,
		CONTEXT_MENU, SLEEP, SELECT, EXECUTE, HELP,
	}

	for _, k := range keysToCheck {
		t.Run(fmt.Sprintf("Key_0x%X", uint32(k)), func(t *testing.T) {
			name, exists := keyMap[k]
			if !exists {
				t.Errorf("Key 0x%X is not defined in keyMap!", uint32(k))
			}
			if name == "" {
				t.Errorf("Key 0x%X is defined as empty string!", uint32(k))
			}
		})
	}
}

func TestGetKey(t *testing.T) {
	key, ok := getKey(uint32(A))
	if !ok || key != A {
		t.Errorf("getKey returned incorrect result for key A")
	}

	// Test phím tào lao
	_, ok = getKey(0xFFF)
	if ok {
		t.Errorf("getKey should return false for unknown key code")
	}
}
