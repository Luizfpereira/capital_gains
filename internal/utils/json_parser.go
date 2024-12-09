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
	var operations [][]domain.Operation
	res := strings.Split(input, "][")
	for i := range res {
		if len(res) > 1 {
			if i%2 == 0 {
				res[i] += "]"
			} else {
				res[i] = "[" + res[i]
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
