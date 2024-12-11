package adapter

import (
	"capital_gains/internal/domain"
	"encoding/json"
	"fmt"
)

type Writer interface {
	WriteTaxes(taxes [][]domain.Tax)
}

type ConsoleWriter struct{}

func NewConsoleWriter() *ConsoleWriter {
	return &ConsoleWriter{}
}

func (w *ConsoleWriter) WriteTaxes(taxes [][]domain.Tax) {
	for _, taxList := range taxes {
		result, err := json.Marshal(taxList)
		if err != nil {
			fmt.Printf("Error marshalling taxes: %v\n", err)
		}
		fmt.Println(string(result))
	}
}
