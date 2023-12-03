package matrix

import (
	"fmt"
)

type Matrix struct {
	Data [][]byte
}

func New() *Matrix {
	return &Matrix{Data: [][]byte{}}
}

func (m *Matrix) Read(line string) error {
	l := []byte{}

	for _, c := range line {
		l = append(l, byte(c))
	}

	m.Data = append(m.Data, l)

	return nil
}

func (m *Matrix) IsEmpty(i, j int) bool {
	return m.Data[i][j] == '.'
}

func (m *Matrix) Log(logger func(format string, args ...any)) {
	if len(m.Data) == 0 {
		return
	}

	// Align the matrix line numbers nicely - determine necessary max width in digits.
	digits := len(fmt.Sprintf("%d", len(m.Data)-1))
	df := fmt.Sprintf("%%%dd", digits)

	for i, l := range m.Data {
		logger(df+": %s", i, l)
	}
}

func (m *Matrix) Walk(ii Indexes, proc func(i, j int) bool) {
	for _, i := range ii {
		ok := proc(i.I, i.J)
		if !ok {
			return
		}
	}
}
