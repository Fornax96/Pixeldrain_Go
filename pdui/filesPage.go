package pdui

import (
	"fmt"

	"github.com/visualfc/goqt/ui"
)

func (w *MainWindow) setPageFiles() {
	fmt.Println("Files page button pressed")
	w.clearContent()

	fmw := NewFileManagerWidget()

	l := ui.NewHBoxLayout()
	l.SetMargin(0)
	l.AddWidget(fmw)

	w.contentArea.SetLayout(l)

}
