package settings

import (
	"encoding/json"
	"hidtool/app/logger"
	"hidtool/app/util"
	"os"
	"sync"
)

const fileName = "settings.json"

// SettingsImpl mirrors settings.json. Pointers keep unset keys distinguishable
// from an explicit false, so a missing key falls back to its default.
type SettingsImpl struct {
	RunOnStartup *bool `json:"run_on_startup,omitempty"`
}

var (
	mutex   sync.Mutex
	current SettingsImpl
)

func path() string {
	folder := ""
	if util.IsDebug() {
		folder = "build"
	}
	return util.GetPath(folder, fileName)
}

func init() {
	data, err := os.ReadFile(path())
	if err != nil {
		if !os.IsNotExist(err) {
			logger.Warn("Failed to read settings:", err)
		}
		return
	}

	if err := json.Unmarshal(data, &current); err != nil {
		logger.Warn("Failed to parse settings, using defaults:", err)
		current = SettingsImpl{}
	}
}

func save() error {
	data, err := json.MarshalIndent(current, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path(), append(data, '\n'), 0644)
}

// RunOnStartup reports whether the app should register itself to start with
// Windows. Defaults to true when the user has never chosen.
func RunOnStartup() bool {
	mutex.Lock()
	defer mutex.Unlock()

	if current.RunOnStartup == nil {
		return true
	}
	return *current.RunOnStartup
}

func SetRunOnStartup(enabled bool) error {
	mutex.Lock()
	defer mutex.Unlock()

	current.RunOnStartup = &enabled
	return save()
}
