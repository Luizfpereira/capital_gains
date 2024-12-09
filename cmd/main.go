package main

import (
	"capital_gains/internal/adapter"
	"capital_gains/internal/usecase"
	"capital_gains/internal/utils"
	"log"
)

func main() {
	reader := adapter.NewStdinReader()
	input, err := reader.ReadInput()
	if err != nil {
		log.Fatalf("Failed to read input: %v", err)
	}

	jsonParser := utils.NewJSONParser()
	operations, err := jsonParser.ParseOperations(input)
	if err != nil {
		log.Fatalf("Failed to parse operations: %v", err)
	}

	operationProcessor := usecase.NewOperationProcessor()
	taxes := operationProcessor.ProcessOperations(operations)

	writer := adapter.NewConsoleWriter()
	writer.WriteTaxes(taxes)
}
