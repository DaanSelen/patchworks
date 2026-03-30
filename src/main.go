package main

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	draw "patchworks/src/modules/draw"
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

	userEntry := widget.NewEntry()
	passEntry := widget.NewPasswordEntry()
	targEntry := widget.NewEntry()
	var book string
	inputContainer := draw.MakeInputContainer(userEntry, passEntry, targEntry, &book, w)

	content := container.NewBorder(
		inputContainer, // top
		nil,            // bottom
		nil,            // left
		nil,            // right
		nil,            // center
	)

	w.SetContent(content)
	w.ShowAndRun()
}
