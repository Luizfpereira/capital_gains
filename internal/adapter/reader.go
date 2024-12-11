package adapter

import (
	"bufio"
	"os"
	"strings"
)

type Reader interface {
	ReadInput() (string, error)
}

type StdinReader struct{}

func NewStdinReader() *StdinReader {
	return &StdinReader{}
}

func (r *StdinReader) ReadInput() (string, error) {
	scanner := bufio.NewScanner(os.Stdin)
	var input []string
	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "" {
			break
		}
		input = append(input, line)
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}
	return strings.Join(input, ""), nil
}
