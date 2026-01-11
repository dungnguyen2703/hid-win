package logger

import (
	"hidtool/app/util"
	"log"
)

// create DEBUG logger that can be enabled or disabled via build tags

func Debug(v ...any) {
	if util.IsDebug() {
		log.Println("[DEBUG]", v)
	}
}

func Info(v ...any) {
	if util.IsDebug() {
		log.Println("[INFO]", v)
	}
}

func Warn(v ...any) {
	if util.IsDebug() {
		log.Println("[WARN]", v)
	}
}

func Error(v ...any) {
	if util.IsDebug() {
		log.Println("[ERROR]", v)
	}
}
