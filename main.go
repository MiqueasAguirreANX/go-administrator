package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	"github.com/MiqueasAguirreANX/GOAdministrator/stores"
)

func main() {

	itemStore := stores.ItemStore{}
	itemStore.InitStore()

	a := app.New()
	w := a.NewWindow("GO-Administrator")

	w.Resize(fyne.NewSize(1024, 720))

	w.SetContent(widget.NewLabel("This my go implementation of the administrator package"))
	w.ShowAndRun()
}
