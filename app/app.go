package app

import (
	"hidtool/app/event"
	"hidtool/app/keyboard"
	"hidtool/app/mice"
	"hidtool/app/profile"
	"os"
	"syscall"
	"unsafe"

	"github.com/getlantern/systray"
	"golang.org/x/sys/windows/registry"
)

var (
	user32           = syscall.NewLazyDLL("user32.dll")
	setWindowsHookEx = user32.NewProc("SetWindowsHookExW")
	callNextHookEx   = user32.NewProc("CallNextHookEx")
	getMessage       = user32.NewProc("GetMessageW")
)

const (
	WH_KEYBOARD_LL = 13
	WH_MOUSE_LL    = 14
)

func debugLog(v ...any) {
	// log.Println(v...)
}

func run() {
	hook, _, _ := setWindowsHookEx.Call(
		WH_MOUSE_LL,
		syscall.NewCallback(func(nCode int, wParam uintptr, lParam uintptr) uintptr {
			currentProfile := profile.GetProfile()
			if currentProfile == profile.All || currentProfile == profile.Mice {
				key, ok := mice.Check(nCode, wParam, lParam)
				if ok {
					debugLog("🐭 Mouse Event:", key)
					switch key {
					case mice.BACK_BUTTON_DOWN:
						event.Run(event.WindowLeft)
						return 1
					case mice.FORWARD_BUTTON_DOWN:
						event.Run(event.WindowRight)
						return 1
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
			key, modifiers, ok := keyboard.Check(nCode, wParam, lParam)
			if ok {
				debugLog("⌨️ Keyboard Event:", key, modifiers)
				if len(modifiers) == 0 {
					currentProfile := profile.GetProfile()
					if currentProfile == profile.All || currentProfile == profile.Keyboard {
						switch key {
						case keyboard.F1:
							event.Run(event.WindowLeft)
							return 1
						case keyboard.F2:
							event.Run(event.WindowRight)
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

func RunAtStartup() error {
	execPath, err := os.Executable()
	if err != nil {
		return err
	}

	key, _, err := registry.CreateKey(registry.CURRENT_USER, `Software\Microsoft\Windows\CurrentVersion\Run`, registry.SET_VALUE)
	if err != nil {
		return err
	}
	defer key.Close()

	return key.SetStringValue("HIDTool", execPath)
}

func RunInSysTray(iconData []byte) {
	// Init Systray
	systray.SetIcon(iconData)
	systray.SetTitle("HID Tool")
	systray.SetTooltip("HID Tool - Mouse and Keyboard Enhancer")

	// Initiaize Menu
	mOff := systray.AddMenuItem("Off", "Disable all hooks")
	mAll := systray.AddMenuItem("Mice + Keyboard", "Enable both")
	mKey := systray.AddMenuItem("Keyboard Only", "Enable keyboard only")
	mMice := systray.AddMenuItem("Mice Only", "Enable mice only")
	systray.AddSeparator()
	mQuit := systray.AddMenuItem("Quit", "Quit the application")

	menuMap := map[profile.Profile]*systray.MenuItem{
		profile.Off:      mOff,
		profile.All:      mAll,
		profile.Keyboard: mKey,
		profile.Mice:     mMice,
	}

	updateCheckmarks := func() {
		current := profile.GetProfile()
		for name, item := range menuMap {
			if name == current {
				item.Check()
			} else {
				item.Uncheck()
			}
		}
	}

	updateCheckmarks()

	// Run main app logic
	go run()

	// Loop to handle menu item clicks
	go func() {
		for {
			select {
			case <-mOff.ClickedCh:
				profile.SetProfile(profile.Off)
				updateCheckmarks()
			case <-mAll.ClickedCh:
				profile.SetProfile(profile.All)
				updateCheckmarks()
			case <-mKey.ClickedCh:
				profile.SetProfile(profile.Keyboard)
				updateCheckmarks()
			case <-mMice.ClickedCh:
				profile.SetProfile(profile.Mice)
				updateCheckmarks()
			case <-mQuit.ClickedCh:
				systray.Quit()
			}
		}
	}()

}
