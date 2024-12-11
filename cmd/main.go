package main

import (
	"capital_gains/cmd/app"
	"capital_gains/internal/adapter"
	"capital_gains/internal/usecase"
	"capital_gains/internal/utils"
	"log"
)

func main() {
	reader := adapter.NewStdinReader()
	parser := utils.NewJSONParser()
	processor := usecase.NewOperationProcessor()
	writer := adapter.NewConsoleWriter()

	if err := app.Run(reader, parser, processor, writer); err != nil {
		log.Fatalf("Application error: %v", err)
	}
}
