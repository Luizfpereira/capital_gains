package app

import (
	"capital_gains/internal/adapter"
	"capital_gains/internal/usecase"
	"capital_gains/internal/utils"
	"fmt"
)

func Run(reader adapter.Reader, parser *utils.JSONParser, processor *usecase.OperationProcessor, writer adapter.Writer) error {
	input, err := reader.ReadInput()
	if err != nil {
		return fmt.Errorf("failed to read input: %w", err)
	}

	operations, err := parser.ParseOperations(input)
	if err != nil {
		return fmt.Errorf("failed to parse operations: %w", err)
	}

	taxes := processor.ProcessOperations(operations)
	writer.WriteTaxes(taxes)
	return nil
}
