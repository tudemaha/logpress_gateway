package routes

import (
	"log"
	"net/http"

	compressController "github.com/tudemaha/logpress_gateway/internal/compress/controller"
	pingController "github.com/tudemaha/logpress_gateway/internal/ping"
)

func LoadRoutes() {
	log.Println("INFO LoadRoutes: loading routes...")

	http.HandleFunc("/ping", pingController.PingHandler())
	http.HandleFunc("/compress", compressController.CompressHandler())

	log.Println("INFO LoadRoutes: routes loaded.")
}
