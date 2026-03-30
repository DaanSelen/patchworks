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

//go:embed Icon.png
var iconData []byte

func main() {
	if ok, err := setup.PrepareEnvironment(); !ok {
		log.Printf("An error occurred: %v", err)
		return
	}

	iconResource := fyne.NewStaticResource("./src/Icon.png", iconData)

	app := app.NewWithID("nl.nerthus.patchworks")
	app.Settings().SetTheme(theme.DefaultTheme())
	app.SetIcon(iconResource)

	w := app.NewWindow("PatchWorks")
	w.Resize(windowSize)

	userEntry := widget.NewEntry()
	passEntry := widget.NewPasswordEntry()
	targEntry := widget.NewEntry()
	var book string
	
	inputContainer := draw.MakeInputContainer(userEntry, passEntry, targEntry, &book, w)
	infoContainer := draw.MakeInfoContainer()

	content := container.NewBorder(
		nil, // top
		nil,            // bottom
		nil,            // left
		nil,            // right
		inputContainer,            // center
	)

	w.SetContent(content)
	w.ShowAndRun()
}
