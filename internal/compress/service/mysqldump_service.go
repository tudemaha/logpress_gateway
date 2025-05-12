package service

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/tudemaha/logpress_gateway/pkg/logpress"
)

func CreateDump() (string, error) {
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	tableName := os.Getenv("TABLE_NAME")

	output, _ := exec.Command("pwd").Output()
	pwd := strings.TrimSpace(string(output))
	id := uuid.New().String()

	now := time.Now()

	dumpArgs := []string{
		"--add-drop-database=false",
		"--add-drop-table=false",
		"--no-create-db=true",
		"--no-create-info=true",
		"--single-transaction",
		"-u", username,
		"-p" + password,
		dbName, tableName,
		fmt.Sprintf("--result-file=%s/%s/%s.sql", pwd, "dump/uncompressed", id),
	}
	exec.Command("mysqldump", dumpArgs...).Output()
	time.Sleep(5 * time.Second)

	logpressConfig := logpress.LoadLogpressConfig
	logpressConfig.LastDumpTimestamp = now

	if err := logpress.WriteConfig(logpressConfig); err != nil {
		return "", err
	}

	logpress.ReadConfig()

	return id, nil
}
