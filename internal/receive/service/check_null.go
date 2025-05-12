package service

import "github.com/tudemaha/logpress_gateway/internal/receive/dto"

func CheckNullValue(data *dto.NewSensorData) bool {
	if data.NodeID == "" ||
		data.Temp == 0 ||
		data.Humid == 0 ||
		data.SoilPH == 0 ||
		data.SoilMoisture == 0 ||
		data.Gas == 0 ||
		data.GPS == "" {
		return true
	}
	return false
}
