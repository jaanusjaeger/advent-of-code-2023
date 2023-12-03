package main

import "github.com/jaanusjaeger/advent-of-code-2023/log"

type processor struct {
	sum int
}

func newProcessor() *processor {
	return &processor{}
}

func (p *processor) process(line string) error {
	if line == "" {
		return nil
	}

	first := digit(line, 0, 1)
	last := digit(line, len(line)-1, -1)

	log.Debugf("line %q, first %d, last %d", line, first, last)

	num := first*10 + last

	p.sum += num

	return nil
}

func (p *processor) result() (any, error) {
	return p.sum, nil
}

func digit(s string, start, inc int) int {
	// No end condition - expect the input to always contain a digit
	for i := start; ; i += inc {
		c := s[i]
		if c >= '0' && c <= '9' {
			return int(c - '0')
		}
	}
}
