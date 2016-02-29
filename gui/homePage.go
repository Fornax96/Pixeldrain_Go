// homePage
package gui

import (
	"fmt"
	"strconv"

	"github.com/visualfc/goqt/ui"
)

func (w *MainWindow) setPageHome() {
	fmt.Println("Home Button Pressed")
	w.clearContent()

	l := ui.NewVBoxLayout()
	l.SetAlignment(ui.Qt_AlignTop)

	for i := 0; i < 20; i++ {
		obj1 := ui.NewPushButton()
		obj1.SetMinimumHeight(50)
		obj1.SetText("btn " + strconv.Itoa(i))

		l.AddWidget(obj1)
	}

	w.contentArea.SetLayout(l)
}
