package usecase

import "capital_gains/internal/utils"

func CalculateWeightedAverage(oldAverage, newCost float64, oldQuantity, newQuantity int) float64 {
	totalQuantity := float64(oldQuantity + newQuantity)
	return utils.Round(((float64(oldQuantity) * oldAverage) + (float64(newQuantity) * newCost)) / totalQuantity)
}
