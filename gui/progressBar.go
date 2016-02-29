package gui

import (
	"github.com/visualfc/goqt/ui"
)

type ProgressBar struct {
	*ui.QProgressBar
	window *ui.QWidget
}

func NewProgressBar() *ProgressBar {
	pb := &ProgressBar{
		QProgressBar: ui.NewProgressBar(),
		window:       ui.NewWidget(),
	}

	return pb
}

func (pb *ProgressBar) Show() {
	pb.window.Show()
}

func (pb *ProgressBar) SetProgress() {
	pb.SetValue(50)
}

func (pb *ProgressBar) Destroy() {
	pb.window.Destroy(true, true)
}
