package service

import (
	"math"

	"github.com/tudemaha/logpress_gateway/internal/receive/dto"
)

func TransformFloat(value float64, precision uint8) float64 {
	pow := math.Pow(10, float64(precision))
	transformed := math.Round(value*pow) / pow

	return transformed
}

func TransformSensorData(data *dto.NewSensorData) {
	data.Temp = TransformFloat(data.Temp, 14)
	data.Humid = TransformFloat(data.Humid, 14)
	data.SoilPH = TransformFloat(data.SoilPH, 14)
	data.SoilMoisture = TransformFloat(data.SoilMoisture, 14)
	data.Gas = TransformFloat(data.Gas, 14)
}
