package gui

import (
	"fmt"

	"github.com/visualfc/goqt/ui"
)

type MainWindow struct {
	*ui.QWidget
	application *ui.QApplication

	toolbar       *ui.QVBoxLayout
	btnHome       *ui.QPushButton
	btnFiles      *ui.QPushButton
	btnUploadText *ui.QPushButton
	btnSettings   *ui.QPushButton
	scrollArea    *ui.QScrollArea
	contentArea   *ui.QWidget

	trayMenu *TrayMenu
}

func (w *MainWindow) clearContent() {
	children := w.contentArea.Children()

	for _, c := range children {
		c.Delete()
	}
}

func (w *MainWindow) OnCloseEvent(ev *ui.QCloseEvent) bool {
	fmt.Println("Close Event Caught")
	w.Hide()

	ev.Ignore()

	return false
}

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

func NewMainWindow(args []string) *MainWindow {
	w := &MainWindow{}

	w.application = ui.NewApplication(args)
	w.application.SetQuitOnLastWindowClosed(false)

	fmt.Printf("Available styles: %v \n", ui.NewStyleFactory().Keys())
	w.SetTheme("DarkFusion")

	w.btnHome = ui.NewPushButton()
	w.btnHome.SetText("Home")
	w.btnHome.OnClicked(w.setPageHome)

	w.btnFiles = ui.NewPushButton()
	w.btnFiles.SetText("My Files")
	w.btnFiles.OnClicked(w.setPageFiles)

	w.btnUploadText = ui.NewPushButton()
	w.btnUploadText.SetText("Upload Text")
	w.btnUploadText.OnClicked(w.setPageUploadText)

	w.btnSettings = ui.NewPushButton()
	w.btnSettings.SetText("Settings")
	w.btnSettings.OnClicked(w.setPageSettings)

	w.toolbar = ui.NewVBoxLayout()
	w.toolbar.SetAlignment(ui.Qt_AlignTop)
	w.toolbar.AddWidget(w.btnHome)
	w.toolbar.AddWidget(w.btnFiles)
	w.toolbar.AddWidget(w.btnUploadText)
	w.toolbar.AddWidget(w.btnSettings)

	w.contentArea = ui.NewWidget()

	w.scrollArea = ui.NewScrollArea()
	w.scrollArea.SetWidgetResizable(true) // Make the scrollbar appear
	w.scrollArea.SetWidget(w.contentArea)
	w.scrollArea.SetHorizontalScrollBarPolicy(ui.Qt_ScrollBarAlwaysOff)

	// Main HBox to divide the toolbar buttons and the content area
	hbox := ui.NewHBoxLayout()
	hbox.AddLayout(w.toolbar)
	hbox.AddWidget(w.scrollArea)

	// Create window
	w.QWidget = ui.NewWidget()
	w.SetWindowTitle("Pixeldrain_Go")
	w.SetMinimumHeight(300)
	w.SetMinimumWidth(500)
	w.SetLayout(hbox)
	// The scrollarea makes the window bigger, so we just change it back here
	w.ResizeWithWidthHeight(500, 300)

	// Enable event listening
	w.InstallEventFilter(w)

	// Load up the Home page so we don't show an empty window
	w.setPageHome()

	// Init the tray icon
	w.trayMenu = NewTrayMenu(w)

	return w
}
