package service

import (
	"github.com/tudemaha/logpress_gateway/pkg/database"
	"github.com/tudemaha/logpress_gateway/pkg/logpress"
)

func DeleteOldData() error {
	db := database.DatabaseConnection()
	defer db.Close()

	logpressConfig := logpress.LoadLogpressConfig

	stmt := "DELETE FROM sensors WHERE timestamp < ?"
	_, err := db.Exec(stmt, logpressConfig.LastDumpTimestamp)
	if err != nil {
		return err
	}

	return nil
}
