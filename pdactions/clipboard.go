package pdactions

import (
	"fmt"

	"github.com/visualfc/goqt/ui"
)

func ClipboardCopy(text string) {
	if text == "" {
		return
	}

	//	url := ui.NewUrlWithUrl(text)
	//	desktopServices := ui.NewDesktopServices()
	//	desktopServices.OpenUrl(url)
	//desktopServices.Delete()

	clip := ui.Application().Clipboard()

	clip.SetText(text)

	clip.Delete()

	fmt.Println("Copied to clipboard: " + text)
}
