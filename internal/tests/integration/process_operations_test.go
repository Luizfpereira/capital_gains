package integration

import (
	"capital_gains/cmd/app"
	"capital_gains/internal/adapter/mocks"
	"capital_gains/internal/usecase"
	"capital_gains/internal/utils"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRunIntegration(t *testing.T) {

	tests := []struct {
		name           string
		mockInput      string
		expectedOutput string
	}{
		{
			name: "sale below 20000",
			mockInput: `[{"operation":"buy","unit-cost":10.0,"quantity":3000}]
		[{"operation":"sell","unit-cost":15.0,"quantity":2000}]`,
			expectedOutput: `[{"tax":0}]
		[{"tax":2000}]`,
		},
		{
			name: "multiple sales below 20000",
			mockInput: `[{
						"operation":"buy", "unit-cost":10.00, "quantity": 100},
						{"operation":"sell", "unit-cost":15.00, "quantity": 50},
						{"operation":"sell", "unit-cost":15.00, "quantity": 50}
					]`,
			expectedOutput: `[{"tax":0},{"tax":0},{"tax":0}]`,
		},
		{
			name: "sales with profit and loss",
			mockInput: `[{
						"operation":"buy", "unit-cost":10.00, "quantity": 10000},
						{"operation":"sell", "unit-cost":20.00, "quantity": 5000},
						{"operation":"sell", "unit-cost":5.00, "quantity": 5000}
					]`,
			expectedOutput: `[{"tax":0},{"tax":10000},{"tax":0}]`,
		},
		{
			name: "combination of cases",
			mockInput: `[{"operation":"buy", "unit-cost":10.00, "quantity": 100},
								{"operation":"sell", "unit-cost":15.00, "quantity": 50},
								{"operation":"sell", "unit-cost":15.00, "quantity": 50}]
								[{"operation":"buy", "unit-cost":10.00, "quantity": 10000},
								{"operation":"sell", "unit-cost":20.00, "quantity": 5000},
								{"operation":"sell", "unit-cost":5.00, "quantity": 5000}]`,
			expectedOutput: `[{"tax":0},{"tax":0},{"tax":0}]
		[{"tax":0},{"tax":10000},{"tax":0}]`,
		},
		{
			name: "loss and profit (with deduction)",
			mockInput: `[{"operation":"buy", "unit-cost":10.00, "quantity": 10000},
		{"operation":"sell", "unit-cost":5.00, "quantity": 5000},
		{"operation":"sell", "unit-cost":20.00, "quantity": 3000}]`,
			expectedOutput: `[{"tax":0},{"tax":0},{"tax":1000}]`,
		},
		{
			name: "neither profit nor loss",
			mockInput: `[{"operation":"buy", "unit-cost":10.00, "quantity": 10000},
		{"operation":"buy", "unit-cost":25.00, "quantity": 5000},
		{"operation":"sell", "unit-cost":15.00, "quantity": 10000}]`,
			expectedOutput: `[{"tax":0},{"tax":0},{"tax":0}]`,
		},
		{
			name: "neither profit nor loss and profit after",
			mockInput: `[{"operation":"buy", "unit-cost":10.00, "quantity": 10000},
		{"operation":"buy", "unit-cost":25.00, "quantity": 5000},
		{"operation":"sell", "unit-cost":15.00, "quantity": 10000},
		{"operation":"sell", "unit-cost":25.00, "quantity": 5000}]`,
			expectedOutput: `[{"tax":0},{"tax":0},{"tax":0},{"tax":10000}]`,
		},
		{
			name: "one loss, two profits with deduction with zero taxes and a profit with tax",
			mockInput: `[{"operation":"buy", "unit-cost":10.00, "quantity": 10000},
{"operation":"sell", "unit-cost":2.00, "quantity": 5000},
{"operation":"sell", "unit-cost":20.00, "quantity": 2000},
{"operation":"sell", "unit-cost":20.00, "quantity": 2000},
{"operation":"sell", "unit-cost":25.00, "quantity": 1000}]`,
			expectedOutput: `[{"tax":0},{"tax":0},{"tax":0},{"tax":0},{"tax":3000}]`,
		},
		{
			name: "complex operation",
			mockInput: `[{"operation":"buy", "unit-cost":10.00, "quantity": 10000},
{"operation":"sell", "unit-cost":2.00, "quantity": 5000},
{"operation":"sell", "unit-cost":20.00, "quantity": 2000},
{"operation":"sell", "unit-cost":20.00, "quantity": 2000},
{"operation":"sell", "unit-cost":25.00, "quantity": 1000},
{"operation":"buy", "unit-cost":20.00, "quantity": 10000},
{"operation":"sell", "unit-cost":15.00, "quantity": 5000},
{"operation":"sell", "unit-cost":30.00, "quantity": 4350},
{"operation":"sell", "unit-cost":30.00, "quantity": 650}]`,
			expectedOutput: `[{"tax":0},{"tax":0},{"tax":0},{"tax":0},{"tax":3000},{"tax":0},{"tax":0},{"tax":3700},{"tax":0}]`,
		},
		{
			name: "high value taxes",
			mockInput: `[{"operation":"buy", "unit-cost":10.00, "quantity": 10000},
{"operation":"sell", "unit-cost":50.00, "quantity": 10000},
{"operation":"buy", "unit-cost":20.00, "quantity": 10000},
{"operation":"sell", "unit-cost":50.00, "quantity": 10000}]`,
			expectedOutput: `[{"tax":0},{"tax":80000},{"tax":0},{"tax":60000}]`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockReader := &mocks.MockReader{Input: tt.mockInput, Err: nil}
			var writtenData []string
			mockWriter := &mocks.MockWriter{WrittenData: writtenData}

			parser := utils.NewJSONParser()
			processor := usecase.NewOperationProcessor()

			err := app.Run(mockReader, parser, processor, mockWriter)
			assert.NoError(t, err)

			expected := strings.ReplaceAll(tt.expectedOutput, "\t", "")

			output := strings.Join(mockWriter.WrittenData, "\n")
			assert.Equal(t, expected, output)
		})
	}
}
