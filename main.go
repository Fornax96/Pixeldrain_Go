package main

import (
	"fmt"
	"os"

	"github.com/Fornax96/pixeldrain-go/pdui"
	"github.com/visualfc/goqt/ui"
)

func main() {
	fmt.Println("Ayy")

	ui.RunEx(os.Args, func() {
		pdui.NewMainWindow().Show()
	})

	fmt.Println("Lmao")
}
