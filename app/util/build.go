package util

import (
	"os"
	"strings"
)

var isDebug bool = false

func InitializeDebugBuild() error {
	execPath, err := os.Executable()
	if err != nil {
		return err
	}
	isDebug = strings.Contains(execPath, "AppData\\Local\\Temp")
	return nil
}

func IsDebugBuild() bool {
	return isDebug
}
