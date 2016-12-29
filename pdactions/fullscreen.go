package pdactions

import (
	"fmt"
	"time"

	//	"github.com/visualfc/goqt/ui"
)

//type FullScreenUpload struct {
//	//	window *MainWindow
//}

//func (f *FullScreenUpload) FullScreen() {
func FullScreen() {

	time.Sleep(time.Second) // Wait a second for the menu to disappear

	file, err := GrabScreen()
	if err != nil {
		fmt.Println("lolwat")
	}

	id, err := NewFileUploader().UploadFile(file)

	if err != nil {
		fmt.Println(err)
	}

	ClipboardCopy(id)

	//	f.window.TrayMenu.ShowMessage("Message Title", "Message content", ui.QSystemTrayIcon_Information, 1000)
}
