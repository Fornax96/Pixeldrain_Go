package pdui

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
	btnExit       *ui.QPushButton
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

func NewMainWindow() *MainWindow {
	w := &MainWindow{}

	w.application = ui.Application()
	//w.application = ui.NewApplication(args) // This was a mistake
	w.application.SetQuitOnLastWindowClosed(false)
	w.application.SetAutoGC(true)

	fmt.Printf("Available styles: %v \n", ui.NewStyleFactory().Keys())
	//w.SetTheme("DarkFusion")

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

	w.btnExit = ui.NewPushButton()
	w.btnExit.SetText("Exit")
	w.btnExit.OnClicked(w.application.Exit)

	w.toolbar = ui.NewVBoxLayout()
	w.toolbar.SetAlignment(ui.Qt_AlignTop)
	w.toolbar.AddWidget(w.btnHome)
	w.toolbar.AddWidget(w.btnFiles)
	w.toolbar.AddWidget(w.btnUploadText)
	w.toolbar.AddWidget(w.btnSettings)
	w.toolbar.AddWidget(w.btnExit)

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
