package main

import (
	"github.com/jaanusjaeger/advent-of-code-2023/log"
	"github.com/jaanusjaeger/advent-of-code-2023/matrix"
	"github.com/jaanusjaeger/advent-of-code-2023/str"
)

type processor struct {
	m *matrix.Matrix
}

func newProcessor() *processor {
	return &processor{
		m: matrix.New(),
	}
}

func (p *processor) process(line string) error {
	return p.m.Read(line)
}

func (p *processor) result() (any, error) {
	p.m.Log(log.Debugf)

	var partNumberSum int

	isSymbol := func(i, j int) bool {
		return !p.m.IsEmpty(i, j) && !str.IsDigit(p.m.Data[i][j])
	}

	for i, l := range p.m.Data {
		for j := 0; j < len(l); j++ {
			digits := p.digitStreak(i, j)

			log.Debugf("digits: %v (Empty: %t)", digits, digits.Empty())

			if digits.Empty() {
				continue
			}

			j += digits.SizeJ()

			// Check adjacent.
			adj := digits.Adjacent()
			var hasSymbol bool
			p.m.Walk(adj, func(i, j int) bool {
				if isSymbol(i, j) {
					hasSymbol = true
					return false
				}
				return true
			})

			// If has adjacent symbol, take value of the digits and aggregate.
			if hasSymbol {
				val := str.MustParseInt(string(digits.Value()))
				partNumberSum += val
			}
		}
	}

	return partNumberSum, nil
}

func (p *processor) digitStreak(i, j int) matrix.Sub {
	j2 := j
	for ; j2 < len(p.m.Data[i]) && str.IsDigit(p.m.Data[i][j2]); j2++ {
	}
	return matrix.Sub{Matrix: p.m, I1: i, I2: i, J1: j, J2: j2 - 1}
}
