package service

import (
	"bytes"
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

func TransferCompressedDump(filename string) error {
	serverUrl := os.Getenv("SERVER_UPLOAD")

	file, err := os.Open("./dump/compressed/" + filename + ".sql.gz")
	if err != nil {
		return err
	}
	defer file.Close()

	var b bytes.Buffer
	writer := multipart.NewWriter(&b)

	fw, err := writer.CreateFormFile("file", filename+".sql.gz")
	if err != nil {
		return err
	}

	if _, err = io.Copy(fw, file); err != nil {
		return err
	}

	if err := writer.Close(); err != nil {
		return err
	}

	req, err := http.NewRequest("POST", serverUrl, &b)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)
	if res.StatusCode != 200 {
		return errors.New(string(body))
	}

	return nil
}
