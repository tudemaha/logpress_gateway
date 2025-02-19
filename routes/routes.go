package routes

import (
	"log"
	"net/http"

	authController "github.com/tudemaha/logpress_gateway/internal/auth/controller"
	compressController "github.com/tudemaha/logpress_gateway/internal/compress/controller"
	pingController "github.com/tudemaha/logpress_gateway/internal/ping"
)

func LoadRoutes() {
	log.Println("INFO LoadRoutes: loading routes...")

	fs := http.FileServer(http.Dir("public"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/login", authController.LoginHandler())

	http.HandleFunc("/ping", pingController.PingHandler())
	http.HandleFunc("/compress", compressController.CompressHandler())

	log.Println("INFO LoadRoutes: routes loaded.")
}
