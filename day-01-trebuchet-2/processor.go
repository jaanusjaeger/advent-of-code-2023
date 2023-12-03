package main

import (
	"strings"

	"github.com/jaanusjaeger/advent-of-code-2023/log"
)

var search map[string]int = map[string]int{
	"0":     0,
	"1":     1,
	"2":     2,
	"3":     3,
	"4":     4,
	"5":     5,
	"6":     6,
	"7":     7,
	"8":     8,
	"9":     9,
	"zero":  0,
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

type processor struct {
	sum int
}

func (p *processor) process(line string) error {
	if line == "" {
		return nil
	}

	first := find(line, strings.Index, lt)
	last := find(line, strings.LastIndex, gt)

	log.Debugf("line %q, first %d, last %d", line, first, last)

	num := first*10 + last

	p.sum += num

	return nil
}

func (p *processor) result() (any, error) {
	return p.sum, nil
}

func find(s string, finder func(s, substring string) int, cmp func(i, j int) bool) int {
	res := -1
	aggInd := -1
	for key, val := range search {
		i := finder(s, key)
		if i >= 0 {
			if aggInd < 0 || cmp(i, aggInd) {
				res = val
				aggInd = i
			}
		}
	}
	return res
}

func lt(i, j int) bool {
	return i < j
}

func gt(i, j int) bool {
	return i > j
}
