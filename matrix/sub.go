package matrix

import "fmt"

type Sub struct {
	*Matrix
	I1, I2 int
	J1, J2 int
}

func (s Sub) SizeI() int {
	return s.I2 - s.I1 + 1
}

func (s Sub) SizeJ() int {
	return s.J2 - s.J1 + 1
}

func (s Sub) Size() int {
	return s.SizeI() * s.SizeJ()
}

func (s Sub) Empty() bool {
	return s.Size() == 0
}

func (s Sub) Adjacent() Indexes {
	if s.Empty() {
		return nil
	}

	ii := []Index{}
	minJ := max(0, s.J1-1)
	maxJ := min(len(s.Data[s.I1])-1, s.J2+1)

	// Preceding row.
	if s.I1 > 0 {
		for j := minJ; j <= maxJ; j++ {
			ii = append(ii, Index{I: s.I1 - 1, J: j})
		}
	}
	// Same rows.
	for i := s.I1; i <= s.I2; i++ {
		if minJ < s.J1 {
			ii = append(ii, Index{I: i, J: minJ})
		}
		if maxJ > s.J2 {
			ii = append(ii, Index{I: i, J: maxJ})
		}
	}
	// Following row.
	if s.I2 < len(s.Data)-1 {
		for j := minJ; j <= maxJ; j++ {
			ii = append(ii, Index{I: s.I2 + 1, J: j})
		}
	}

	return Indexes(ii)
}

func (s Sub) Value() []byte {
	res := []byte{}

	for i := s.I1; i <= s.I2; i++ {
		res = append(res, s.Data[i][s.J1:s.J2+1]...)
	}

	return res
}

func (s Sub) String() string {
	return fmt.Sprintf("M: %p, i: {%d, %d}, j: {%d, %d}", s.Matrix, s.I1, s.I2, s.J1, s.J2)
}
