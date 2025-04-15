package service

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/google/uuid"
)

func CreateDump() string {
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	name := os.Getenv("DB_NAME")

	output, _ := exec.Command("pwd").Output()
	pwd := strings.TrimSpace(string(output))
	id := uuid.New().String()

	dumpArgs := []string{
		"--add-drop-database=false",
		"--add-drop-table=false",
		"--no-create-db=true",
		"--no-create-info=true",
		"--single-transaction",
		"-u", username,
		"-p" + password,
		name, "sensors",
		fmt.Sprintf("--result-file=%s/%s/%s.sql", pwd, "dump/uncompressed", id),
	}
	exec.Command("mysqldump", dumpArgs...).Output()
	time.Sleep(5 * time.Second)

	return id
}
