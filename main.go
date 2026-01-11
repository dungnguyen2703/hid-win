package main

import (
	_ "embed"
	"hidtool/app"
	"hidtool/app/logger"
	"hidtool/app/util"

	"github.com/getlantern/systray"
)

//go:embed icon.ico
var iconData []byte

func main() {
	go app.Run()
	if !util.IsDebug() {
		app.RunOnStartup()
	}
	systray.Run(onReady, onExit)
}

func onReady() {
	app.SetupSysTray(iconData)
	logger.Info("Application Started")
}

func onExit() {
	// Clean up here
}
