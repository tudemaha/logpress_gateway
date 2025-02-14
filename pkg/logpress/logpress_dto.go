package logpress

import "time"

type LogpressConfig struct {
	Username          string    `json:"username"`
	Password          string    `json:"password"`
	Threshold         uint8     `json:"threshold"`
	ThresholdUnit     string    `json:"threshold_unit"`
	LastDumpTimestamp time.Time `json:"last_dump_timestamp"`
}
