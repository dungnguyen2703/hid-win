package main

import (
	_ "embed"
	"hidtool/app"
	"hidtool/app/logger"
	"hidtool/app/profile"
	"hidtool/app/util"
	"os"
	"syscall"
	"time"
	"unsafe"

	"github.com/getlantern/systray"
	"golang.org/x/sys/windows/registry"
)

//go:embed icon.ico
var iconData []byte

var (
	kernel32        = syscall.NewLazyDLL("kernel32.dll")
	procCreateEvent = kernel32.NewProc("CreateEventW")
	procSetEvent    = kernel32.NewProc("SetEvent")
	procWaitForObj  = kernel32.NewProc("WaitForSingleObject")
	globalHEvent    uintptr
)

func runOnStartup() error {
	execPath, _ := os.Executable()
	key, _, err := registry.CreateKey(registry.CURRENT_USER, `Software\Microsoft\Windows\CurrentVersion\Run`, registry.SET_VALUE)
	if err != nil {
		return err
	}
	defer key.Close()

	return key.SetStringValue("HIDTool", execPath)
}

func setupSysTray() {
	systray.SetIcon(iconData)
	systray.SetTitle("HID Tool")
	systray.SetTooltip("HID Tool - Mouse and Keyboard Enhancer")
	menuMap := map[profile.Profile]*systray.MenuItem{}
	for _, profile := range profile.List {
		menu := systray.AddMenuItem(profile.GetName(), profile.GetDescription())
		menuMap[profile] = menu
	}
	systray.AddSeparator()
	mQuit := systray.AddMenuItem("Quit", "Quit the application")

	updateCheckmarks := func() {
		current := profile.GetCurrentProfile()
		for pf, mItem := range menuMap {
			if current != nil && pf.GetID() == current.GetID() {
				mItem.Check()
			} else {
				mItem.Uncheck()
			}
		}
	}

	updateCheckmarks()

	go func() {
		for range mQuit.ClickedCh {
			systray.Quit()
			return
		}
	}()

	for pf, mItem := range menuMap {
		go func(p profile.Profile, item *systray.MenuItem) {
			for range item.ClickedCh {
				current := profile.GetCurrentProfile()
				if (current == nil) || (current.GetID() != p.GetID()) {
					profile.SetCurrentProfile(p)
				} else {
					profile.SetCurrentProfile(nil)
				}
				updateCheckmarks()
			}
		}(pf, mItem)
	}

}

func check() {
	eventName, _ := syscall.UTF16PtrFromString("Local\\HidTool_Unique_Instance")
	h, _, err := procCreateEvent.Call(0, 0, 0, uintptr(unsafe.Pointer(eventName)))
	if h == 0 {
		return
	}
	globalHEvent = h
	if err != nil && err.(syscall.Errno) == 183 {
		procSetEvent.Call(globalHEvent)
		time.Sleep(500 * time.Millisecond)
		h, _, _ = procCreateEvent.Call(0, 0, 0, uintptr(unsafe.Pointer(eventName)))
		globalHEvent = h
	}
	go func() {
		for {
			ret, _, _ := procWaitForObj.Call(globalHEvent, 0xFFFFFFFF)
			if ret == 0 {
				logger.Info("New instance detected, exiting...")
				systray.Quit()
				return
			}
		}
	}()
}

func main() {
	check()
	go app.Run()
	if !util.IsDebug() {
		runOnStartup()
	}
	systray.Run(onReady, onExit)
}

func onReady() {
	setupSysTray()
	logger.Info("Application Started")
}

func onExit() {
	logger.Info("Application Exited")
}
