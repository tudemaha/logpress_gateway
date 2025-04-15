package logpress

import "time"

type LogpressConfig struct {
	Username          string    `json:"username"`
	Password          string    `json:"password"`
	Threshold         uint16    `json:"threshold"`
	ThresholdUnit     string    `json:"threshold_unit"`
	LastDumpTimestamp time.Time `json:"last_dump_timestamp"`
	CronInterval      uint8     `json:"cron_interval"`
	CronUnit          string    `json:"cron_unit"`
}
