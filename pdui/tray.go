package pdui

import (
	"fmt"

	"fornax96.me/pixeldrain/pdactions"

	"github.com/visualfc/goqt/ui"
)

type TrayMenu struct {
	*ui.QSystemTrayIcon
	parentWindow *MainWindow
	icon         *ui.QIcon
	menu         *ui.QMenu

	menuToggleWindow      *ui.QAction
	menuScreenshot        *ui.QAction
	menuPartialScreenshot *ui.QAction
	menuEditScreenshot    *ui.QAction
	menuCloseApp          *ui.QAction
}

func (t *TrayMenu) toggleWindow() {
	if t.parentWindow.IsVisible() {
		t.parentWindow.Hide()
		t.ShowMessage(
			"Window hidden",                               // Title
			"Pixeldrain is now running in the background", // Content
			ui.QSystemTrayIcon_Information,                // Icon
			2000, // Time
		)
	} else {
		t.parentWindow.Show()
	}
}

func (t *TrayMenu) closeApp() {
	t.parentWindow.application.Exit()
}

func (t *TrayMenu) prepareMenu() {
	if t.parentWindow.IsVisible() {
		t.menuToggleWindow.SetText("Hide Window")
	} else {
		t.menuToggleWindow.SetText("Show Window")
	}
}

func (t *TrayMenu) activated(reason ui.QSystemTrayIcon_ActivationReason) {
	fmt.Println("activated")

	if reason == ui.QSystemTrayIcon_MiddleClick {
		fmt.Println("middle click")
	}
}

func NewTrayMenu(parent *MainWindow) *TrayMenu {
	t := &TrayMenu{
		QSystemTrayIcon: ui.NewSystemTrayIcon(),
		parentWindow:    parent,
		icon:            ui.NewIcon(),
		menu:            ui.NewMenu(),
	}

	t.icon.AddFile("res/pixeldrain_transparent.png")

	t.SetIcon(t.icon)
	t.OnActivated(t.activated)
	t.Show()

	//t.ShowMessage("Message Title", "Message content", ui.QSystemTrayIcon_Information, 1000)

	t.menuToggleWindow = ui.NewAction(t.menu)
	t.menuToggleWindow.SetText("Hide Window")
	t.menuToggleWindow.OnTriggered(t.toggleWindow)

	t.menuScreenshot = ui.NewAction(t.menu)
	t.menuScreenshot.SetText("Capture Screen")
	t.menuScreenshot.OnTriggered(func() {
		go pdactions.FullScreen()
		fmt.Println("Capture Screen")
	})

	t.menuPartialScreenshot = ui.NewAction(t.menu)
	t.menuPartialScreenshot.SetText("Capture Part Of Screen")

	t.menuEditScreenshot = ui.NewAction(t.menu)
	t.menuEditScreenshot.SetText("Open Editor")

	t.menuCloseApp = ui.NewAction(t.menu)
	t.menuCloseApp.SetText("Quit Pixeldrain")
	t.menuCloseApp.OnTriggered(t.closeApp)

	// Add the menu items to the menu
	t.menu.AddAction(t.menuToggleWindow)
	t.menu.AddSeparator()
	t.menu.AddAction(t.menuScreenshot)
	t.menu.AddAction(t.menuPartialScreenshot)
	t.menu.AddAction(t.menuEditScreenshot)
	t.menu.AddSeparator()
	t.menu.AddAction(t.menuCloseApp)

	t.menu.OnAboutToShow(t.prepareMenu)

	t.SetContextMenu(t.menu)

	return t
}
