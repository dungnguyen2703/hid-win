package main

import (
	_ "embed"
	"hidtool/app"
	"hidtool/app/util"

	"github.com/getlantern/systray"
)

//go:embed icon.ico
var iconData []byte

func main() {
	err := util.InitializeDebugBuild()
	if err != nil {
		panic(err)
	}
	go app.Run()
	if !util.IsDebugBuild() {
		app.RunOnStartup()
	}
	systray.Run(onReady, onExit)
}

func onReady() {
	app.SetupSysTray(iconData)
}

func onExit() {
	// Clean up here
}
