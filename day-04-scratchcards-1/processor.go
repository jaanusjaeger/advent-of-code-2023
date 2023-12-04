package main

import (
	"github.com/jaanusjaeger/advent-of-code-2023/log"
	"github.com/jaanusjaeger/advent-of-code-2023/set"
	"github.com/jaanusjaeger/advent-of-code-2023/str"
)

type processor struct {
	score int
}

func newProcessor() *processor {
	return &processor{}
}

func (p *processor) process(line string) error {
	if line == "" {
		return nil
	}

	cardRow := str.MustSplit(line, ":", 2)
	numbersRow := str.MustSplit(cardRow[1], "|", 2)
	winningsSlice := str.Split(numbersRow[0], " ")
	selection := str.Split(numbersRow[1], " ")

	winnings := set.New[string]().Add(winningsSlice...)

	var score int
	for _, s := range selection {
		if winnings.Contains(s) {
			if score == 0 {
				score = 1
				continue
			}
			score *= 2
		}
	}

	log.Debugf("line: %s", line)
	log.Debugf("    winning: %s", winnings.ToSlice())
	log.Debugf("      score: %d", score)

	p.score += score

	return nil
}

func (p *processor) result() (any, error) {
	return p.score, nil
}
