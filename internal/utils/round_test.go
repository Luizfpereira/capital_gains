package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRound(t *testing.T) {
	tests := []struct {
		name     string
		input    float64
		expected float64
	}{
		{
			name:     "Round positive number",
			input:    16.6666,
			expected: 16.67,
		},
		{
			name:     "Round negative number",
			input:    -16.6666,
			expected: -16.67,
		},
		{
			name:     "Round exact two decimal places",
			input:    20.33,
			expected: 20.33,
		},
		{
			name:     "Round with trailing zeros",
			input:    20.3300000,
			expected: 20.33,
		},
		{
			name:     "Round up",
			input:    15.255,
			expected: 15.26,
		},
		{
			name:     "Round down",
			input:    15.254,
			expected: 15.25,
		},
		{
			name:     "Round small positive number",
			input:    0.005,
			expected: 0.01,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			rounded := Round(test.input)
			assert.Equal(t, test.expected, rounded)
		})
	}
}
