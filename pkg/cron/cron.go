package cron

import (
	"database/sql"
	"log"
	"os"
	"time"

	"github.com/tudemaha/logpress_gateway/pkg/database"
	"github.com/tudemaha/logpress_gateway/pkg/logpress"
)

func StartCron() {
	log.Println("INFO StartCron: strting cron job...")

	db := database.Connection
	config, err := logpress.ReadConfig()
	if err != nil {
		log.Fatalf("ERROR StartCron fatal error: %v", err)
	}

	log.Println("INFO StartCron: cron job started.")

	var dbSize float64
	var rows *sql.Rows
	for {
		stmt := "SELECT ROUND(SUM(data_length + index_length) / 1024 / 1024"
		if config.ThresholdUnit == "GB" {
			stmt += " / 1024"
		}
		stmt += ", 1)"
		stmt += " FROM information_schema.TABLES"
		stmt += " WHERE table_schema = ?"
		stmt += " GROUP BY table_schema"

		rows, err = db.Query(stmt, os.Getenv("DB_NAME"))
		if err != nil {
			log.Fatalf("ERROR cron job fatal error: %v", err)
		}

		if rows.Next() {
			_ = rows.Scan(&dbSize)
		}

		log.Println(dbSize)

		if config.CronUnit == "sec" {
			time.Sleep(time.Duration(config.CronInterval) * time.Second)
		} else {
			time.Sleep(time.Duration(config.CronInterval) * time.Minute)
		}
	}
}
