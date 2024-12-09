package usecase

import (
	"capital_gains/internal/domain"
	"capital_gains/internal/utils"
)

type OperationProcessor struct{}

func NewOperationProcessor() *OperationProcessor {
	return &OperationProcessor{}
}

func (p *OperationProcessor) ProcessOperations(operations [][]domain.Operation) [][]domain.Tax {
	var weightedAverage, loss float32
	var totalQuantity int
	var taxes [][]domain.Tax

	for _, operationsList := range operations {
		var taxList []domain.Tax
		for _, operation := range operationsList {
			var tax float32

			switch operation.Operation {
			case domain.BUY:
				weightedAverage = CalculateWeightedAverage(weightedAverage, totalQuantity, operation.UnitCost, operation.Quantity)
				totalQuantity += operation.Quantity
			case domain.SELL:
				totalQuantity -= operation.Quantity
				salesPrice := utils.Round(float32(operation.Quantity) * operation.UnitCost)
				averagePrice := utils.Round(float32(operation.Quantity) * weightedAverage)

				tax, loss = CalculateTaxAndLoss(salesPrice, averagePrice, loss)
			}
			taxList = append(taxList, domain.Tax{Tax: tax})
		}
		taxes = append(taxes, taxList)
	}
	return taxes
}
