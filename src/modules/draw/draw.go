package draw

import (
	"image/color"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
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
		canvas.NewLine(color.Gray{Y: 128}),
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
		log.Printf("updated selected book to: %s", *book)
	}
	bookBox := container.NewBorder(
		widget.NewLabel("Available Books"), // top
		nil, nil, nil,                      // bottom, left, right
		bookList, // center fills
	)

	inputBox := container.NewBorder(
		targetBox,
		nil, nil, nil,
		bookBox,
	)
	return inputBox
}

func MakeFooter(targEntry *widget.Entry, book *string, app fyne.App) *fyne.Container {
	var result string
	textEntry := widget.NewMultiLineEntry()

	showBtn := widget.NewButton("Show Results", func() {
		w := app.NewWindow("Result Display")
		w.Resize(windowSize)

		w.SetContent(container.NewScroll(textEntry))
		w.Show()
	})
	showWrap := container.NewGridWrap(
		buttonSize,
		showBtn,
	)
	showBtn.Disable()

	var actionBtn *widget.Button
	actionBtn = widget.NewButton("Execute", func() {
		showBtn.Disable()

		actionBtn.Importance = widget.HighImportance
		actionBtn.Refresh()

		log.Println("beginning execution with external binary")

		ok, path := runner.FindMeshbookBinary()
		if !ok {
			log.Println("something went wrong while looking for the binary, see above for details")
		}

		log.Println("kicking off goroutine")
		actionBtn.Disable()
		go func() {
			ok, result = runner.RunMeshbook(path, *book, targEntry.Text)
			if !ok {
				log.Println("assuming failed state")
				fyne.CurrentApp().Driver().DoFromGoroutine(func() {
					actionBtn.Importance = widget.DangerImportance
					actionBtn.Refresh()
				}, true)

			} else {
				fyne.CurrentApp().Driver().DoFromGoroutine(func() {
					textEntry.SetText(result)

					actionBtn.Importance = widget.SuccessImportance
					actionBtn.Refresh()

					showBtn.Enable()
				}, true)
			}

			fyne.CurrentApp().Driver().DoFromGoroutine(func() {
				actionBtn.Enable()
			}, true)

			if len(result) > 0 {
				log.Println(result)
			}
		}()

	})
	actionWrap := container.NewGridWrap(
		buttonSize,
		actionBtn,
	)

	cancelBtn := widget.NewButton("Exit", func() {
		log.Println("quitting")
		app.Quit()
	})
	cancelWrap := container.NewGridWrap(
		buttonSize,
		cancelBtn,
	)

	bottomBox := container.NewHBox(
		actionWrap, // left
		showWrap,
		//layout.NewSpacer(), // flexible space
		cancelWrap, // right
	)
	return bottomBox
}
