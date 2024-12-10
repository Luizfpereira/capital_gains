package usecase

import (
	"capital_gains/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProcessOperations(t *testing.T) {
	processor := NewOperationProcessor()

	tests := []struct {
		name       string
		operations [][]domain.Operation
		expected   [][]domain.Tax
	}{
		{
			name: "Single buy operation",
			operations: [][]domain.Operation{
				{
					{Operation: domain.BUY, UnitCost: 10.0, Quantity: 5000},
				},
			},
			expected: [][]domain.Tax{
				{
					{Tax: 0},
				},
			},
		},
		{
			name: "Single sell operation with profit",
			operations: [][]domain.Operation{
				{
					{Operation: domain.BUY, UnitCost: 10.0, Quantity: 10000},
					{Operation: domain.SELL, UnitCost: 15.0, Quantity: 5000},
				},
			},
			expected: [][]domain.Tax{
				{
					{Tax: 0},
					{Tax: 5000},
				},
			},
		},
		{
			name: "Single sell operation with loss",
			operations: [][]domain.Operation{
				{
					{Operation: domain.BUY, UnitCost: 10.0, Quantity: 10000},
					{Operation: domain.SELL, UnitCost: 5.0, Quantity: 10000},
				},
			},
			expected: [][]domain.Tax{
				{
					{Tax: 0},
					{Tax: 0},
				},
			},
		},
		{
			name: "Sell operation with exemption (sales under 20000)",
			operations: [][]domain.Operation{
				{
					{Operation: domain.BUY, UnitCost: 10.0, Quantity: 100},
					{Operation: domain.SELL, UnitCost: 50.0, Quantity: 100},
				},
			},
			expected: [][]domain.Tax{
				{
					{Tax: 0},
					{Tax: 0},
				},
			},
		},
		{
			name: "Multiple operations",
			operations: [][]domain.Operation{
				{
					{Operation: domain.BUY, UnitCost: 10.0, Quantity: 1000},
					{Operation: domain.BUY, UnitCost: 20.0, Quantity: 5000},
					{Operation: domain.SELL, UnitCost: 15.0, Quantity: 500},
					{Operation: domain.SELL, UnitCost: 30.0, Quantity: 2000},
				},
			},

			expected: [][]domain.Tax{
				{
					{Tax: 0},
					{Tax: 0},
					{Tax: 0},
					{Tax: 4335},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := processor.ProcessOperations(tt.operations)
			assert.Equal(t, tt.expected, result)
		})
	}
}
