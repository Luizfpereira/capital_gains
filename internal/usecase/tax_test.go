package usecase

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateTaxAndLoss(t *testing.T) {
	tests := []struct {
		name         string
		salesPrice   float64
		averagePrice float64
		oldLoss      float64
		expectedTax  float64
		expectedLoss float64
	}{
		{
			name:         "Loss scenario",
			salesPrice:   5000,
			averagePrice: 7000,
			oldLoss:      0,
			expectedTax:  0,
			expectedLoss: 2000,
		},
		{
			name:         "Profit without exceeding price threshold",
			salesPrice:   15000,
			averagePrice: 10000,
			oldLoss:      0,
			expectedTax:  0,
			expectedLoss: 0,
		},
		{
			name:         "Profit exceeding price threshold with no loss",
			salesPrice:   25000,
			averagePrice: 10000,
			oldLoss:      0,
			expectedTax:  3000,
			expectedLoss: 0,
		},
		{
			name:         "Profit exceeding price threshold with loss",
			salesPrice:   25000,
			averagePrice: 20000,
			oldLoss:      5000,
			expectedTax:  0,
			expectedLoss: 0,
		},
		{
			name:         "No profit but old loss",
			salesPrice:   20000,
			averagePrice: 20000,
			oldLoss:      3000,
			expectedTax:  0,
			expectedLoss: 3000,
		},
		{
			name:         "Profit exceeding price threshold fully covering loss",
			salesPrice:   30000,
			averagePrice: 20000,
			oldLoss:      3000,
			expectedTax:  1400,
			expectedLoss: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tax, loss := CalculateTaxAndLoss(tt.salesPrice, tt.averagePrice, tt.oldLoss)
			assert.Equal(t, tt.expectedTax, tax, "Tax should match")
			assert.Equal(t, tt.expectedLoss, loss, "Loss should match")
		})
	}
}
