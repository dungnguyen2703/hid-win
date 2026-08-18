package startup

import (
	"errors"
	"hidtool/app/logger"
	"hidtool/app/util"
	"os"

	"golang.org/x/sys/windows/registry"
)

const (
	runKeyPath = `Software\Microsoft\Windows\CurrentVersion\Run`
	valueName  = "HIDTool"
)

// Apply syncs the Windows Run key with the desired state. Debug builds leave
// the registry untouched, so running from source never registers itself.
func Apply(enabled bool) error {
	if util.IsDebug() {
		logger.Debug("Skipping startup registration in debug build, enabled:", enabled)
		return nil
	}
	if enabled {
		return enable()
	}
	return disable()
}

// enable writes the current executable path, so moving the app keeps it valid
func enable() error {
	execPath, err := os.Executable()
	if err != nil {
		return err
	}

	key, _, err := registry.CreateKey(registry.CURRENT_USER, runKeyPath, registry.SET_VALUE)
	if err != nil {
		return err
	}
	defer key.Close()

	return key.SetStringValue(valueName, execPath)
}

func disable() error {
	key, err := registry.OpenKey(registry.CURRENT_USER, runKeyPath, registry.SET_VALUE)
	if err != nil {
		if errors.Is(err, registry.ErrNotExist) {
			return nil
		}
		return err
	}
	defer key.Close()

	if err := key.DeleteValue(valueName); err != nil && !errors.Is(err, registry.ErrNotExist) {
		return err
	}
	return nil
}
