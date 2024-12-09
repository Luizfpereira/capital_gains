package usecase

import "capital_gains/internal/utils"

func CalculateWeightedAverage(oldAverage float32, oldQuantity int, newCost float32, newQuantity int) float32 {
	totalQuantity := float32(oldQuantity + newQuantity)
	return utils.Round(((float32(oldQuantity) * oldAverage) + (float32(newQuantity) * newCost)) / totalQuantity)
}
