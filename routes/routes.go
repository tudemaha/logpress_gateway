package routes

import (
	"log"
	"net/http"

	pingController "github.com/tudemaha/logpress_gateway/internal/ping"
)

func LoadRoutes() {
	log.Println("INFO LoadRoutes: loading routes...")

	http.HandleFunc("/ping", pingController.PingHandler())

	log.Println("INFO LoadRoutes: routes loaded.")
}
