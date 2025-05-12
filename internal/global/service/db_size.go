package service

import (
	"fmt"
	"os"

	"github.com/tudemaha/logpress_gateway/pkg/database"
	"github.com/tudemaha/logpress_gateway/pkg/logpress"
)

func GetDBSize() (float64, error) {
	var dbSize float64
	config := logpress.LoadLogpressConfig
	db := database.DatabaseConnection()
	defer db.Close()

	stmt := "SELECT ROUND(SUM(data_length + index_length) / 1024 / 1024"
	if config.ThresholdUnit == "GB" {
		stmt += " / 1024"
	}
	stmt += ", 1)"
	stmt += " FROM information_schema.TABLES"
	stmt += " WHERE table_schema = ?"
	stmt += " GROUP BY table_schema"

	_, err := db.Exec(fmt.Sprintf("ANALYZE TABLE %s.%s", os.Getenv("DB_NAME"), os.Getenv("TABLE_NAME")))
	if err != nil {
		return dbSize, err
	}
	rows, err := db.Query(stmt, os.Getenv("DB_NAME"))
	if err != nil {
		return dbSize, err
	}

	if rows.Next() {
		_ = rows.Scan(&dbSize)
	}

	return dbSize, nil
}
