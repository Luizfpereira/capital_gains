package utils

import "math"

func Round(x float32) float32 {
	return float32(math.Round(float64(x)*100) / 100)
}
