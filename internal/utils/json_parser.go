package utils

import (
	"capital_gains/internal/domain"
	"encoding/json"
	"errors"
	"strings"
)

type JSONParser struct{}

func NewJSONParser() *JSONParser {
	return &JSONParser{}
}

func (p *JSONParser) ParseOperations(input string) ([][]domain.Operation, error) {
	normalizedInput := strings.ReplaceAll(input, "\n", "")
	normalizedInput = strings.ReplaceAll(normalizedInput, "\t", "")
	normalizedInput = strings.ReplaceAll(normalizedInput, " ", "")

	var operations [][]domain.Operation
	res := strings.Split(normalizedInput, "][")
	for i := range res {
		if len(res) > 1 {
			if i == 0 {
				res[i] += "]"
			} else if i == len(res)-1 {
				res[i] = "[" + res[i]
			} else {
				res[i] = "[" + res[i] + "]"
			}
		}
		var ops []domain.Operation
		err := json.Unmarshal([]byte(res[i]), &ops)
		if err != nil {
			return nil, errors.New("invalid JSON format")
		}
		operations = append(operations, ops)
	}
	return operations, nil
}
