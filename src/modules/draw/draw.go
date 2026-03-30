package draw

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"

	tasks "patchworks/src/modules/tasks"
)

func MakeInputContainer(userEntry, passEntry, targEntry *widget.Entry, book *string, win fyne.Window) *fyne.Container {
	availBooks, err := tasks.ListAvailableBooks()
	if err != nil {
		log.Println("failed to read available books on disk")
	}

	loginBox := container.NewVBox(
		widget.NewLabel("Service Username"),
		userEntry,
		widget.NewLabel("Service Password"),
		passEntry,
	)

	targetBox := container.NewVBox(
		widget.NewLabel("Target Device or Group"),
		targEntry,
	)

	bookList := widget.NewList(
		func() int {
			return len(availBooks)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("Books")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(availBooks[i])
		},
	)
	bookList.OnSelected = func(id widget.ListItemID) {
		*book = availBooks[id]
		log.Printf("Updated selected book to: %s", *book)
	}
	bookBox := container.NewBorder(
		widget.NewLabel("Available Books"), // top
		nil, nil, nil,                      // bottom, left, right
		bookList, // center fills
	)

	inputBox := container.New(
		layout.NewGridLayoutWithColumns(3),
		loginBox,  // column 1
		targetBox, // column 2
		bookBox,   // column 3
	)

	return inputBox
}
