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

	gearIndexToDigits := map[matrix.Index][]matrix.Sub{}

	isGearSymbol := func(i, j int) bool {
		return p.m.Data[i][j] == '*'
	}

	for i, l := range p.m.Data {
		for j := 0; j < len(l); j++ {
			digits := p.digitStreak(i, j)

			log.Debugf("digits: %v (Empty: %t)", digits, digits.Empty())

			if digits.Empty() {
				continue
			}

			j += digits.SizeJ()

			// Collect adjacent gear symbol indexes.
			adj := digits.Adjacent()
			p.m.Walk(adj, func(i, j int) bool {
				if isGearSymbol(i, j) {
					gearIndex := matrix.Index{I: i, J: j}
					gearIndexToDigits[gearIndex] = append(gearIndexToDigits[gearIndex], digits)
				}
				return true
			})
		}
	}

	// Find the gears, i.e. where gear symbol is adjacent to exactly 2 digits.
	var gearRatioSum int
	for _, digits := range gearIndexToDigits {
		if len(digits) != 2 {
			continue
		}
		val1 := str.MustParseInt(string(digits[0].Value()))
		val2 := str.MustParseInt(string(digits[1].Value()))
		gearRatio := val1 * val2
		gearRatioSum += gearRatio
	}

	return gearRatioSum, nil
}

func (p *processor) digitStreak(i, j int) matrix.Sub {
	j2 := j
	for ; j2 < len(p.m.Data[i]) && str.IsDigit(p.m.Data[i][j2]); j2++ {
	}
	return matrix.Sub{Matrix: p.m, I1: i, I2: i, J1: j, J2: j2 - 1}
}
