// fileManager
package pdui

import (
	"strconv"

	"github.com/visualfc/goqt/ui"
)

type FileManagerWidget struct {
	*ui.QListWidget
}

func NewFileManagerWidget() *FileManagerWidget {
	f := &FileManagerWidget{}

	f.QListWidget = ui.NewListWidget()
	f.SetViewMode(ui.QListView_IconMode)
	f.SetFlow(ui.QListView_TopToBottom)
	f.SetResizeMode(ui.QListView_Adjust)
	f.SetDragEnabled(false)
	f.SetSpacing(10)
	f.SetWrapping(true)
	//f.QAbstractScrollArea.SetHorizontalScrollBarPolicy(ui.qsc)
	f.SetSelectionMode(ui.QAbstractItemView_ExtendedSelection)
	f.SetAlternatingRowColors(false)

	for i := 0; i <= 1000; i++ {
		label := ui.NewLabel()
		label.SetText("yoghurt #" + strconv.Itoa(i))

		pix := ui.NewLabel()
		//icon := ui.NewIconWithFilename("/home/wim/Pictures/agario.png")
		pix.SetPixmap(
			ui.NewIconWithFilename("/home/wim/Pictures/1287002805.or.21616.png").Pixmap(
				ui.NewSizeWithWidthHeight(24, 24)))

		layout := ui.NewHBoxLayout()
		layout.SetMargin(0)
		layout.AddWidget(pix)
		layout.AddWidget(label)

		wid := ui.NewWidget()
		wid.SetLayout(layout)

		li := ui.NewListWidgetItem()
		li.SetSizeHint(wid.SizeHint())
		//li.SetSizeHint(ui.NewSizeWithWidthHeight(140, 24))

		f.AddItem(li)
		f.SetItemWidget(li, wid)
	}

	return f
}
