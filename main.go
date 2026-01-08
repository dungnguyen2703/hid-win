package main

import (
	_ "embed"
	"hidtool/app"

	"github.com/getlantern/systray"
)

//go:embed icon.ico
var iconData []byte

func main() {
	// Systray.Run
	systray.Run(onReady, onExit)
}

func onReady() {
	// Init Systray
	app.RunInSysTray(iconData)
	app.RunAtStartup()

}

func onExit() {
	// Clean up here
}
