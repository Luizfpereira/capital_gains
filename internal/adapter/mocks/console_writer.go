package mocks

import (
	"capital_gains/internal/domain"
	"encoding/json"
)

type MockWriter struct {
	WrittenData []string
}

func (m *MockWriter) WriteTaxes(taxes [][]domain.Tax) {
	for _, taxList := range taxes {
		data, _ := json.Marshal(taxList)
		m.WrittenData = append(m.WrittenData, string(data))
	}
}
