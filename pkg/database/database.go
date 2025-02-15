package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var Connection *sql.DB

func DatabaseConnection() {
	log.Println("INFO DatabaseConnection: trying to connect to database...")

	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))

	Connection, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("ERROR DatabaseConnection fatal error: %v", err)
	}

	Connection.SetConnMaxLifetime(3 * time.Minute)
	Connection.SetMaxOpenConns(10)
	Connection.SetMaxIdleConns(10)

	log.Println("INFO DatabaseConnection: database connected.")
}
