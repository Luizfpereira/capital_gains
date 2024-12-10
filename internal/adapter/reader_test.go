package adapter

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadInput(t *testing.T) {
	tests := []struct {
		name      string
		mockInput string
		expected  string
		expectErr bool
	}{
		{
			name:      "Single line input",
			mockInput: "Hello, World!\n",
			expected:  "Hello, World!",
			expectErr: false,
		},
		{
			name: "Single line input with empty line",
			mockInput: `Hello, World!
			
			`,
			expected:  "Hello, World!",
			expectErr: false,
		},
		{
			name: "Multiple lines input",
			mockInput: `Line 1
Line 2
Line 3

`,
			expected:  "Line 1Line 2Line 3",
			expectErr: false,
		},
		{
			name:      "Empty input",
			mockInput: "\n",
			expected:  "",
			expectErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			oldStdin := os.Stdin
			r, w, _ := os.Pipe()
			buf.WriteString(tt.mockInput)
			buf.WriteTo(w)
			os.Stdin = r
			w.Close()

			defer func() {
				os.Stdin = oldStdin
				r.Close()
			}()

			reader := NewStdinReader()
			result, err := reader.ReadInput()

			if tt.expectErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestReadInputErr(t *testing.T) {
	oldStdin := os.Stdin

	r, w, _ := os.Pipe()
	w.Close()
	r.Close() //close to cause error in scanner
	os.Stdin = r
	defer func() { os.Stdin = oldStdin }()

	reader := NewStdinReader()
	_, err := reader.ReadInput()
	assert.Error(t, err)
}
