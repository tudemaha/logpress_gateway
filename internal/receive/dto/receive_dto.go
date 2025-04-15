package dto

import "time"

type SensorData struct {
	Timestamp time.Time `json:"timestamp"`
	DeviceID  string    `json:"device_id"`
	CO        float32   `json:"co"`
	Humid     float32   `json:"humid"`
	Temp      float32   `json:"temp"`
	LPG       float32   `json:"lpg"`
	Smoke     float32   `json:"smoke"`
	Light     bool      `json:"light"`
	Motion    bool      `json:"motion"`
}
