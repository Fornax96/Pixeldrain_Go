package main

import (
	"fmt"
	"os"

	"fornax96.me/pixeldrain/gui"
	"github.com/visualfc/goqt/ui"
)

func main() {
	fmt.Println("Ayy")

	ui.RunEx(os.Args, func() {
		window := gui.NewMainWindow(os.Args)

		window.Show()
	})

	fmt.Println("Lmao")
}
