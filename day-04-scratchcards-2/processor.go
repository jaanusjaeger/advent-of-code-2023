package main

import (
	"fmt"

	"github.com/jaanusjaeger/advent-of-code-2023/log"
	"github.com/jaanusjaeger/advent-of-code-2023/set"
	"github.com/jaanusjaeger/advent-of-code-2023/str"
)

type processor struct {
	cards  int
	copies map[int]int
}

func newProcessor() *processor {
	return &processor{copies: make(map[int]int)}
}

func (p *processor) process(line string) error {
	if line == "" {
		return nil
	}

	cardRow := str.MustSplit(line, ":", 2)
	cardNoRow := str.Split(cardRow[0], " ")
	cardNo := str.MustParseInt(cardNoRow[1])

	numbersRow := str.MustSplit(cardRow[1], "|", 2)
	winningsSlice := str.Split(numbersRow[0], " ")
	selection := str.Split(numbersRow[1], " ")

	winnings := set.New[string]().Add(winningsSlice...)

	copies := p.copies[cardNo] + 1

	var winCount int
	for _, s := range selection {
		if winnings.Contains(s) {
			winCount++
		}
	}

	for i := 0; i < winCount; i++ {
		p.copies[cardNo+1+i] += copies
	}

	copiesStr := ""
	for i := 1; i <= cardNo || p.copies[i] > 0; i++ {
		copiesStr += fmt.Sprintf("%d: %d, ", i, p.copies[i])
	}
	log.Debugf("copies: %s", copiesStr)
	log.Debugf("line: %s", line)
	log.Debugf("       winCount: %d", winCount)
	log.Debugf("         copies: %d", copies)

	p.cards += copies

	return nil
}

func (p *processor) result() (any, error) {
	return p.cards, nil
}
