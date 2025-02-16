package logpress

import (
	"encoding/json"
	"log"
	"os"
)

var LoadLogpressConfig LogpressConfig

func ReadConfig() {
	log.Println("INFO ReadConfig: loading logpress config...")

	file, err := os.Open("./config/logpress.json")
	if err != nil {
		log.Fatalf("ERROR ReadConfig fatal error: %v", err)
	}
	defer file.Close()

	jsonDecoder := json.NewDecoder(file)
	err = jsonDecoder.Decode(&LoadLogpressConfig)
	if err != nil {
		log.Fatalf("ERROR ReadConfig fatal error: %v", err)
	}

	log.Println("INFO ReadConfig: LogPress config loaded.")
}

func WriteConfig(logpressConfig LogpressConfig) error {
	file, err := os.OpenFile("./config/logpress.json", os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	jsonEncoder := json.NewEncoder(file)
	err = jsonEncoder.Encode(logpressConfig)
	if err != nil {
		return err
	}

	return nil
}
