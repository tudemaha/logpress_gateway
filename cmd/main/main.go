package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/tudemaha/logpress_gateway/pkg/server"
	"github.com/tudemaha/logpress_gateway/routes"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Panicf("ERROR load .env: %v", err)
	}
	routes.LoadRoutes()
	server.StartServer()
}
