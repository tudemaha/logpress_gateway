package service

import (
	"encoding/csv"
	"os"
	"strings"
)

func ReadTransferLog() ([][]string, error) {
	file, err := os.Open("./internal/global/log/transfer.log")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}

	return records, nil
}

func AppendTransferLog(transferLog []string) error {
	file, err := os.OpenFile("./internal/global/log/transfer.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	transferLogStr := strings.Join(transferLog, ",")
	transferLogStr += "\n"

	_, err = file.WriteString(transferLogStr)
	if err != nil {
		return err
	}

	return nil
}
