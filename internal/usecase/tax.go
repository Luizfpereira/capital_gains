package usecase

import "capital_gains/internal/utils"

const (
	TAX_RATE        = 0.2
	PRICE_THRESHOLD = 20000
)

func CalculateTaxAndLoss(salesPrice, averagePrice, oldLoss float64) (tax float64, loss float64) {
	if salesPrice < averagePrice {
		loss = oldLoss + averagePrice - salesPrice
	} else {
		rawProfit := salesPrice - averagePrice
		profit := rawProfit - oldLoss
		if profit > 0 && salesPrice > PRICE_THRESHOLD {
			tax = profit * TAX_RATE
		}
		if rawProfit > oldLoss {
			loss = 0.0
		} else {
			loss = oldLoss - rawProfit
		}
	}
	return utils.Round(tax), utils.Round(loss)
}
