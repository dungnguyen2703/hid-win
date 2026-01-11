//go:build release

package util

import (
	"os"
	"path/filepath"
)

func init() {
	exePath, err := os.Executable()
	if err != nil {
		panic("Failed to get executable path: " + err.Error())
	}
	isDebug = false
	exeDirectory = filepath.Dir(exePath)
}
