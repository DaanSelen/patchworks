package main

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"

	setup "patchworks/src/modules/setup"

	_ "embed"
)

//go:embed icon.ico
var iconData []byte

func main() {
	if ok, err := setup.PrepareEnvironment(); !ok {
		log.Printf("An error occurred: %v", err)
		return
	}

	iconResource := fyne.NewStaticResource("./src/icon.ico", iconData)

	app := app.NewWithID("nl.systemec.patchworks")
	app.Settings().SetTheme(theme.DefaultTheme())
	app.SetIcon(iconResource)

	w := app.NewWindow("PatchWorks")
	w.Resize(windowSize)

	w.ShowAndRun()
}
