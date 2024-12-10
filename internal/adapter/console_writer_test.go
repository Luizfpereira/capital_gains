package adapter

import (
	"bytes"
	"capital_gains/internal/domain"
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWriteTaxes(t *testing.T) {
	taxes := [][]domain.Tax{
		{
			{Tax: 0},
			{Tax: 100.50},
		},
		{
			{Tax: 200.75},
			{Tax: 0},
		},
	}

	var buf bytes.Buffer
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	//redirect Stdout to file from pipe to read from r
	os.Stdout = w

	consoleWriter := NewConsoleWriter()
	consoleWriter.WriteTaxes(taxes)

	w.Close()
	os.Stdout = oldStdout
	buf.ReadFrom(r)

	var expectedOutput bytes.Buffer
	for _, taxList := range taxes {
		result, _ := json.Marshal(taxList)
		expectedOutput.WriteString(string(result) + "\n")
	}

	assert.Equal(t, expectedOutput.String(), buf.String())
}
