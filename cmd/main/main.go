package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/tudemaha/logpress_gateway/pkg/cron"
	"github.com/tudemaha/logpress_gateway/pkg/database"
	"github.com/tudemaha/logpress_gateway/pkg/logpress"
	"github.com/tudemaha/logpress_gateway/pkg/server"
	"github.com/tudemaha/logpress_gateway/routes"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Panicf("ERROR load .env: %v", err)
	}

	database.DatabaseConnection()
	logpress.ReadConfig()

	go cron.StartCron()

	routes.LoadRoutes()
	server.StartServer()
}
