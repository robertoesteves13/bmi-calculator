package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/robertoesteves13/bmi-calculator"
)

func main() {
	a := app.New()
	w := a.NewWindow("BMI Calculator")

	w.Resize(fyne.NewSize(640, 480))

	weight_bind := binding.NewString()
	height_bind := binding.NewString()

	weight_input := widget.NewEntryWithData(weight_bind);
	weight_input.PlaceHolder = "Weight (in kilos)"

	height_input := widget.NewEntryWithData(height_bind);
	height_input.PlaceHolder = "Height (in meters)"

	button_submit := widget.NewButton("Calcular", func() {
		weight_str, err := weight_bind.Get()
		if err != nil {
			dialog.NewError(fmt.Errorf("error when accessing weight bind: %v", err), w).Show()
			return
		}

		height_str, err := height_bind.Get()
		if err != nil {
			dialog.NewError(fmt.Errorf("error when accessing height bind: %v", err), w).Show()
			return
		}

		weight, err := bmicalculator.ParseFloat(weight_str)
		if err != nil {
			dialog.NewError(fmt.Errorf("error parsing weight string: %v", err), w).Show()
			return
		}

		height, err := bmicalculator.ParseFloat(height_str)
		if err != nil {
			dialog.NewError(fmt.Errorf("error parsing height string: %v", err), w).Show()
			return
		}

		if !bmicalculator.IsValidWeight(weight) {
			dialog.NewError(fmt.Errorf("weigh of %f isn't valid", weight), w).Show()
			return
		}

		if !bmicalculator.IsValidHeight(height) {
			dialog.NewError(fmt.Errorf("heigh of %f isn't valid", height), w).Show()
			return
		}
		
		bmi := bmicalculator.CalculateBMI(weight, height)
		dialog.NewInformation("BMI", fmt.Sprintf("BMI is %.2f", bmi), w).Show()
	})

	container := container.NewVBox(
		weight_input,
		height_input,
		layout.NewSpacer(),
		button_submit,
	);
	w.SetContent(container)
	w.ShowAndRun()
}

