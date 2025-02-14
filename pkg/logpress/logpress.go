package logpress

import (
	"encoding/json"
	"os"
)

func ReadConfig() (LogpressConfig, error) {
	var logpressConfig LogpressConfig

	file, err := os.Open("./config/logpress.json")
	if err != nil {
		return logpressConfig, err
	}

	jsonDecoder := json.NewDecoder(file)
	err = jsonDecoder.Decode(&logpressConfig)
	if err != nil {
		return logpressConfig, nil
	}

	return logpressConfig, nil
}

func WriteConfig(logpressConfig LogpressConfig) error {
	file, err := os.OpenFile("./config/logpress.json", os.O_TRUNC|os.O_WRONLY, 0755)
	if err != nil {
		return err
	}

	jsonEncoder := json.NewEncoder(file)
	err = jsonEncoder.Encode(logpressConfig)
	if err != nil {
		return err
	}

	return nil
}
