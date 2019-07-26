package main

import  (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)
func main() {
	app :=app.New()

	w := app.NewWindow("hello")
	w.SetContent(widget.NewVBox(
			widget.NewLabel("hello fyne"),
			widget.NewButton("Quit", func() {
				app.Quit()
			}),
		))
	w.Resize(fyne.NewSize(400,400))
	w.ShowAndRun()
}