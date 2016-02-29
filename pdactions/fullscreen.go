package pdactions

import (
	"fmt"
	"time"
)

func FullScreen() {

	time.Sleep(time.Second) // Wait a second....

	file, err := GrabScreen()
	if err != nil {
		fmt.Println("lolwat")
	}

	id, _ := NewFileUploader().UploadFile(file)

	ClipboardCopy(id)
}
