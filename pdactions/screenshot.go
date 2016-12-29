package pdactions

import (
	"fmt"
	"image/png"
	"os"
	"time"

	"github.com/vova616/screenshot"
)

func GrabScreen() (path string, err error) {
	img, err := screenshot.CaptureScreen()
	if err != nil {
		return "", err
	}

	t := time.Now()

	fname := t.Format("2006-01-02-15-04-05") // Year-Month-Day-Hour-Minute-Second
	fpath := os.TempDir() + "/pd-" + fname + ".png"

	fmt.Println("Screenshot saved as: " + fpath)

	file, err := os.Create(fpath)

	defer file.Close()

	if err != nil {
		return "", err
	}

	png.Encode(file, img)

	img = nil

	return fpath, nil
}
