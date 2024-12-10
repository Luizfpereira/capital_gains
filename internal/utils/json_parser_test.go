package utils

import (
	"capital_gains/internal/domain"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseOperations(t *testing.T) {
	jsonParser := NewJSONParser()

	t.Run("Valid input with multiple nested lists", func(t *testing.T) {
		input := `[{"operation":"buy","unit-cost":10.0,"quantity":100}][{"operation":"sell","unit-cost":15.0,"quantity":50}]`
		expected := [][]domain.Operation{
			{
				{Operation: "buy", UnitCost: 10.0, Quantity: 100},
			},
			{
				{Operation: "sell", UnitCost: 15.0, Quantity: 50},
			},
		}

		result, err := jsonParser.ParseOperations(input)
		assert.NoError(t, err)
		assert.Equal(t, expected, result)
	})

	t.Run("Valid input with multiple lines", func(t *testing.T) {
		input := `[{"operation":"buy","unit-cost":10.0,"quantity":100}]
		[{"operation":"sell","unit-cost":15.0,"quantity":50}]
		[{"operation":"buy","unit-cost":10.0,"quantity":100}]`
		expected := [][]domain.Operation{
			{
				{Operation: "buy", UnitCost: 10.0, Quantity: 100},
			},
			{
				{Operation: "sell", UnitCost: 15.0, Quantity: 50},
			},
			{
				{Operation: "buy", UnitCost: 10.0, Quantity: 100},
			},
		}

		result, err := jsonParser.ParseOperations(input)
		assert.NoError(t, err)
		assert.Equal(t, expected, result)
	})

	t.Run("Invalid input with malformed JSON", func(t *testing.T) {
		input := `[{"operation":"buy","unit-cost":10.0,"quantity":100}, {"operation":`
		result, err := jsonParser.ParseOperations(input)

		assert.Nil(t, result)
		assert.Error(t, err)
		assert.Equal(t, errors.New("invalid JSON format"), err)
	})

	t.Run("Empty input", func(t *testing.T) {
		input := ``

		result, err := jsonParser.ParseOperations(input)
		assert.Nil(t, result)
		assert.Error(t, err)
		assert.Equal(t, errors.New("invalid JSON format"), err)
	})

	t.Run("Single list input", func(t *testing.T) {
		input := `[{"operation":"buy","unit-cost":10.0,"quantity":100}]`
		expected := [][]domain.Operation{
			{
				{Operation: "buy", UnitCost: 10.0, Quantity: 100},
			},
		}

		result, err := jsonParser.ParseOperations(input)
		assert.NoError(t, err)
		assert.Equal(t, expected, result)
	})

	t.Run("Complex input with uneven spacing", func(t *testing.T) {
		input := `[{"operation":"buy","unit-cost":10.0,"quantity":100}] [{"operation":"sell","unit-cost":15.0,"quantity":50}]`
		expected := [][]domain.Operation{
			{
				{Operation: "buy", UnitCost: 10.0, Quantity: 100},
			},
			{
				{Operation: "sell", UnitCost: 15.0, Quantity: 50},
			},
		}

		result, err := jsonParser.ParseOperations(input)
		assert.NoError(t, err)
		assert.Equal(t, expected, result)
	})
}
