package pdactions

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

type UploadedFile struct {
	Id   string
	Type string
	Url  string
}

type FileUploader struct {
}

func NewFileUploader() *FileUploader {
	up := &FileUploader{}

	return up
}

func (fup *FileUploader) UploadFile(uri string) (id string, err error) {
	var b bytes.Buffer
	w := multipart.NewWriter(&b)

	file, err := os.Open(uri)
	if err != nil {
		fmt.Println("Could not open file")
		return "", err
	}

	fw, err := w.CreateFormFile("file", uri)
	if err != nil {
		fmt.Println("Could not create form file")
		return "", err
	}

	io.Copy(fw, file)

	nameField, _ := w.CreateFormField("fileName")
	nameField.Write([]byte(filepath.Base(uri)))

	w.Close()

	// Testing or live version, pick your poison
	req, err := http.NewRequest("POST", "http://213.73.138.240:8080/api/upload", &b)
	//req, err := http.NewRequest("POST", "http://pixeldrain.com/api/upload", &b)
	if err != nil {
		fmt.Println("Could not make request")
		return "", err
	}

	req.Header.Set("Content-Type", w.FormDataContentType())

	client := &http.Client{}
	res, err := client.Do(req)
	defer res.Body.Close()
	if err != nil {
		fmt.Println("Could not do request")
		return "", err
	}

	if res.StatusCode != http.StatusOK {
		err = fmt.Errorf("bad status: %s", res.Status)
		return "", err
	}

	resText, err := ioutil.ReadAll(res.Body)

	response := &UploadedFile{}
	json.Unmarshal(resText, response)

	fmt.Println("File URL: " + response.Url)

	return response.Url, nil
}
