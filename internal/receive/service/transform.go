package service

import "math"

func TransformFloat(value float64, precision uint8) float64 {
	pow := math.Pow(10, float64(precision))
	transformed := math.Round(value*pow) / pow

	return transformed
}
