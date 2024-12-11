package mocks

type MockReader struct {
	Input string
	Err   error
}

func (m *MockReader) ReadInput() (string, error) {
	return m.Input, m.Err
}
