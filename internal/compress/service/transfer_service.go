package service

import (
	"bytes"
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/tudemaha/logpress_gateway/internal/global/dto"
	"github.com/tudemaha/logpress_gateway/internal/global/utils"
)

func TransferCompressedDump(filename string) (dto.ServerResponse, error) {
	serverUrl := os.Getenv("SERVER_UPLOAD") + "/upload"
	var sr dto.ServerResponse

	file, err := os.Open("./dump/compressed/" + filename + ".sql.gz")
	if err != nil {
		return sr, err
	}
	defer file.Close()

	var b bytes.Buffer
	writer := multipart.NewWriter(&b)

	fw, err := writer.CreateFormFile("file", filename+".sql.gz")
	if err != nil {
		return sr, err
	}

	if _, err = io.Copy(fw, file); err != nil {
		return sr, err
	}

	if err := writer.Close(); err != nil {
		return sr, err
	}

	req, err := http.NewRequest("POST", serverUrl, &b)
	if err != nil {
		return sr, err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return sr, err
	}
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)
	if res.StatusCode != 200 {
		return sr, errors.New(string(body))
	}

	sr, err = utils.ParseServerResponse(body)
	if err != nil {
		return sr, err
	}

	return sr, nil
}
