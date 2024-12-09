package usecase

import "capital_gains/internal/utils"

const (
	TAX_RATE        = 0.2
	PRICE_THRESHOLD = 20000
)

func CalculateTaxAndLoss(salesPrice, averagePrice, oldLoss float32) (tax float32, loss float32) {
	if salesPrice < averagePrice {
		loss = oldLoss + averagePrice - salesPrice
	} else {
		rawProfit := salesPrice - averagePrice
		profit := rawProfit - oldLoss
		if profit > 0 && salesPrice > PRICE_THRESHOLD {
			tax = profit * TAX_RATE
		}
		if rawProfit > loss {
			loss = 0.0
		} else {
			loss = oldLoss - rawProfit
		}
	}
	return utils.Round(tax), utils.Round(loss)
}
