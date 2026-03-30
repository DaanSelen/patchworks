package draw

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"

	runner "patchworks/src/modules/runner"
	tasks "patchworks/src/modules/tasks"
)

func MakeInput(targEntry *widget.Entry, book *string) *fyne.Container {
	availBooks, err := tasks.ListAvailableBooks()
	if err != nil {
		log.Println("failed to read available books on disk")
	}

	targetBox := container.NewVBox(
		widget.NewLabel("MeshCentral Target Group"),
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
		layout.NewGridLayoutWithColumns(2),
		targetBox, // column 1
		bookBox,   // column 2
	)
	return inputBox
}

func MakeFooter(targEntry *widget.Entry, book *string, app fyne.App) *fyne.Container {
	actionBtn := widget.NewButton("Execute", func() {
		log.Println("Beginning execution with external binary")

		ok, path := runner.FindMeshbookBinary()
		if !ok {
			log.Println("something went wrong while looking for the binary, see above for details")
		}
		ok, result := runner.RunMeshbook(path, *book, targEntry.Text)
		if !ok {
			log.Println("something went wrong while running the meshbook, see above for details")
		}
		log.Println(result)

	})
	actionWrap := container.NewGridWrap(
		buttonSize,
		actionBtn,
	)

	cancelBtn := widget.NewButton("Exit", func() {
		log.Println("Quitting")
		app.Quit()
	})
	cancelWrap := container.NewGridWrap(
		buttonSize,
		cancelBtn,
	)

	bottomBox := container.NewHBox(
		actionWrap,         // left
		layout.NewSpacer(), // flexible space
		cancelWrap,         // right
	)
	return bottomBox
}
