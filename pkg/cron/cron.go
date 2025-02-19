package cron

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/tudemaha/logpress_gateway/internal/global/service"
	"github.com/tudemaha/logpress_gateway/pkg/logpress"
)

func StartCron() {
	log.Println("INFO StartCron: strting cron job...")

	config := logpress.LoadLogpressConfig
	var dbSize float64
	var err error
	var res *http.Response

	log.Println("INFO StartCron: cron job started.")

	for {
		dbSize, err = service.GetDBSize()
		if err != nil {
			log.Fatalf("ERROR cron job fatal error: %v", err)
		}

		if dbSize > float64(config.Threshold) {
			res, err = http.Post(fmt.Sprintf("http://localhost:%s/compress", os.Getenv("PORT")), "application/json", nil)
			if err != nil {
				log.Fatalf("ERROR cron job fatal error: %v", err)
			}
			log.Printf("INFO cron job compression status code: %d", res.StatusCode)
		}

		switch config.CronUnit {
		case "sec":
			time.Sleep(time.Duration(config.CronInterval) * time.Second)
		case "min":
			time.Sleep(time.Duration(config.CronInterval) * time.Minute)
		case "hour":
			time.Sleep(time.Duration(config.CronInterval) * time.Hour)
		}
	}
}
