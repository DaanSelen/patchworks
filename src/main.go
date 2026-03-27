package main

import (
	"fyne.io/fyne/app"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

//go:embed icon.ico
var iconData []byte

func main() {
	iconResource := fyne.NewStaticResource("./src/icon.ico", iconData)

	app := app.NewWithID("nl.systemec.patchworks")
	app.Settings().SetTheme(theme.DefaultTheme())
	app.SetIcon(iconResource)

}
