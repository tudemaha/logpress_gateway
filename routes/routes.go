package routes

import (
	"log"
	"net/http"

	authController "github.com/tudemaha/logpress_gateway/internal/auth/controller"
	compressController "github.com/tudemaha/logpress_gateway/internal/compress/controller"
	dashboardController "github.com/tudemaha/logpress_gateway/internal/dashboard/controller"
	pingController "github.com/tudemaha/logpress_gateway/internal/ping"
	receiveController "github.com/tudemaha/logpress_gateway/internal/receive/controller"
)

func LoadRoutes() {
	log.Println("INFO LoadRoutes: loading routes...")

	fs := http.FileServer(http.Dir("public"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/login", authController.LoginHandler())
	http.HandleFunc("/", dashboardController.DashboardHandler())
	http.HandleFunc("/config", dashboardController.UpdateConfig())

	http.HandleFunc("/ping", pingController.PingHandler())
	http.HandleFunc("/compress", compressController.CompressHandler())

	http.HandleFunc("/sensors", receiveController.ReceiveHandler())

	log.Println("INFO LoadRoutes: routes loaded.")
}
