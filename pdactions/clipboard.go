package pdactions

import (
	"github.com/visualfc/goqt/ui"
)

func ClipboardCopy(text string) {
	clip := ui.Application().Clipboard()

	clip.SetText(text)
}
