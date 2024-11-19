package main

import (
	"bytes"
	"encoding/json"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("JSON Formatter")

	input := widget.NewMultiLineEntry()
	input.SetPlaceHolder("Write here...")
	input.Wrapping = fyne.TextWrapWord

	output := widget.NewMultiLineEntry()
	output.SetPlaceHolder("Output JSON...")
	output.Wrapping = fyne.TextWrapWord
	output.MultiLine = true
	output.SetText("")

	formatButton := widget.NewButton("Formatter JSON", func() {
		var prettyJSON bytes.Buffer
		err := json.Indent(&prettyJSON, []byte(input.Text), "", "  ")
		if err != nil {
			output.SetText("Error: wrong JSON")
		} else {
			output.SetText(prettyJSON.String())
			output.SetMinRowsVisible(20)
		}
	})

	content := container.NewVBox(
		widget.NewLabel("Input JSON:"),
		container.NewVBox(input),
		formatButton,
		widget.NewLabel("Formatted JSON:"),
		container.NewVBox(output),
	)

	myWindow.SetContent(content)
	myWindow.Resize(fyne.NewSize(800, 600))
	myWindow.ShowAndRun()
}
