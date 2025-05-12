package dto

import "time"

type SensorData struct {
	Timestamp time.Time `json:"timestamp"`
	DeviceID  string    `json:"device_id"`
	CO        float64   `json:"co"`
	Humid     float64   `json:"humid"`
	Temp      float64   `json:"temp"`
	LPG       float64   `json:"lpg"`
	Smoke     float64   `json:"smoke"`
	Light     bool      `json:"light"`
	Motion    bool      `json:"motion"`
}

type NewSensorData struct {
	Timestamp    time.Time `json:"timestamp"`
	NodeID       string    `json:"node_id"`
	Temp         float64   `json:"temp"`
	Humid        float64   `json:"humid"`
	SoilPH       float64   `json:"soil_ph"`
	SoilMoisture float64   `json:"soil_moisture"`
	Gas          float64   `json:"gas"`
	GPS          string    `json:"gps"`
}
