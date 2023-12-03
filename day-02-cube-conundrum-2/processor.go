package main

import (
	"strings"

	"github.com/jaanusjaeger/advent-of-code-2023/log"
	"github.com/jaanusjaeger/advent-of-code-2023/str"
)

type cubes map[string]int

func (c cubes) subset(o cubes) bool {
	for key, val := range c {
		if val > o[key] {
			return false
		}
	}
	return true
}

func (c cubes) power() int {
	if len(c) == 0 {
		return 0
	}

	res := 1

	for _, v := range c {
		res *= v
	}

	return res
}

func max(a, b cubes) cubes {
	res := cubes(map[string]int{})

	max := func(i, j int) int {
		if i > j {
			return i
		}
		return j
	}

	for _, c := range []string{"red", "green", "blue"} {
		res[c] = max(a[c], b[c])
	}

	return res
}

type processor struct {
	sum int
}

func (p *processor) process(line string) error {
	if line == "" {
		return nil
	}

	gg := str.MustSplit(line, ":", 2)

	// Fewest number of cubes, for entire game (over all sets).
	fewest := cubes(map[string]int{})

	sets := strings.Split(gg[1], ";")
	for _, set := range sets {
		cubes := cubes(map[string]int{})

		colorCounts := strings.Split(set, ",")
		for _, cc := range colorCounts {
			ccp := str.MustSplit(strings.TrimSpace(cc), " ", 2)
			cubes[ccp[1]] = str.MustParseInt(ccp[0])
		}

		fewest = max(fewest, cubes)

		log.Debugf("stretch by %v -> %v", cubes, fewest)
	}

	pow := fewest.power()
	log.Debugf("power of cubes %v: %d", fewest, pow)

	p.sum += pow

	return nil
}

func (p *processor) result() (any, error) {
	return p.sum, nil
}
