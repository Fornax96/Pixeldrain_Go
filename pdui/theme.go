// theme
package pdui

import (
	"github.com/visualfc/goqt/ui"
)

func (w *MainWindow) SetTheme(theme string) {
	switch theme {
	case "default":
		w.application.SetStyleWithStyle(ui.NewStyleFactory().Create("GTK+"))

	case "Fusion":
		w.application.SetStyleWithStyle(ui.NewStyleFactory().Create("Fusion"))

	case "DarkFusion":
		w.application.SetStyleWithStyle(ui.NewStyleFactory().Create("Fusion"))

		// This palette was stolen from https://gist.github.com/QuantumCD/6245215
		gray := ui.NewColorWithInt32Int32Int32Int32(53, 53, 53, 255)
		white := ui.NewColorWithName("white")

		palette := ui.NewPalette()
		palette.SetColorWithCrColor(ui.QPalette_Window, gray)
		palette.SetColorWithCrColor(ui.QPalette_WindowText, white)
		palette.SetColorWithCrColor(ui.QPalette_Base, gray)
		palette.SetColorWithCrColor(ui.QPalette_AlternateBase, gray)
		palette.SetColorWithCrColor(ui.QPalette_ToolTipBase, gray)
		palette.SetColorWithCrColor(ui.QPalette_ToolTipText, white)
		palette.SetColorWithCrColor(ui.QPalette_Text, white)
		palette.SetColorWithCrColor(ui.QPalette_Button, gray)
		palette.SetColorWithCrColor(ui.QPalette_ButtonText, white)
		palette.SetColorWithCrColor(ui.QPalette_BrightText, ui.NewColorWithName("red"))
		palette.SetColorWithCrColor(ui.QPalette_Link, ui.NewColorWithInt32Int32Int32Int32(42, 130, 218, 255))

		palette.SetColorWithCrColor(ui.QPalette_Highlight, ui.NewColorWithInt32Int32Int32Int32(42, 130, 218, 255))
		palette.SetColorWithCrColor(ui.QPalette_Highlight, ui.NewColorWithName("black"))

		w.application.SetPalette(palette)
		w.application.SetStyleSheet("QToolTip {border: 1px solid white;}")
	}
}
