package usecase

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateWeightedAverage(t *testing.T) {
	tests := []struct {
		name          string
		oldAverage    float64
		oldQuantity   int
		newCost       float64
		newQuantity   int
		expectedValue float64
	}{
		{
			name:          "Adding to an empty inventory",
			oldAverage:    0,
			oldQuantity:   0,
			newCost:       100,
			newQuantity:   10,
			expectedValue: 100.0,
		},
		{
			name:          "Adding equal cost items",
			oldAverage:    100,
			oldQuantity:   10,
			newCost:       100,
			newQuantity:   10,
			expectedValue: 100.0,
		},
		{
			name:          "Adding higher cost items",
			oldAverage:    100,
			oldQuantity:   10,
			newCost:       150,
			newQuantity:   10,
			expectedValue: 125.0,
		},
		{
			name:          "Adding lower cost items",
			oldAverage:    100,
			oldQuantity:   10,
			newCost:       50,
			newQuantity:   10,
			expectedValue: 75.0,
		},
		{
			name:          "Adding with no new quantity",
			oldAverage:    100,
			oldQuantity:   10,
			newCost:       150,
			newQuantity:   0,
			expectedValue: 100.0,
		},
		{
			name:          "Adding with no old quantity",
			oldAverage:    0,
			oldQuantity:   0,
			newCost:       150,
			newQuantity:   10,
			expectedValue: 150.0,
		},
		{
			name:          "Adding large quantities and costs",
			oldAverage:    1000,
			oldQuantity:   1000,
			newCost:       2000,
			newQuantity:   500,
			expectedValue: 1333.33,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CalculateWeightedAverage(tt.oldAverage, tt.newCost, tt.oldQuantity, tt.newQuantity)
			assert.Equal(t, tt.expectedValue, result)
		})
	}
}
