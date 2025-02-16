package cron

import (
	"log"
	"time"

	"github.com/tudemaha/logpress_gateway/internal/global/service"
	"github.com/tudemaha/logpress_gateway/pkg/logpress"
)

func StartCron() {
	log.Println("INFO StartCron: strting cron job...")

	config := logpress.LoadLogpressConfig
	var dbSize float64
	var err error
	log.Println("INFO StartCron: cron job started.")

	for {
		dbSize, err = service.GetDBSize()
		if err != nil {
			log.Fatalf("ERROR cron job fatal error: %v", err)
		}

		log.Println(dbSize)

		if config.CronUnit == "sec" {
			time.Sleep(time.Duration(config.CronInterval) * time.Second)
		} else {
			time.Sleep(time.Duration(config.CronInterval) * time.Minute)
		}
	}
}
