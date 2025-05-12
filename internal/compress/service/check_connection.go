package service

import (
	"log"
	"net/http"
	"os"
)

func CheckConnection() bool {
	serverUrl := os.Getenv("WAREHOUSE_URL") + "/ping"

	res, err := http.Get(serverUrl)
	if err != nil {
		log.Printf("CheckConnection error: %v", err)
		return false
	}
	res.Body.Close()

	if res.StatusCode == 200 {
		log.Println("CheckConnection: success")
		return true
	}

	return false
}
