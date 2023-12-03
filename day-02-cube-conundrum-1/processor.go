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

var available cubes = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

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

	gg := str.MustSplit(line, ":", 2)
	g := str.MustSplit(gg[0], " ", 2)

	id := str.MustParseInt(g[1])

	sets := strings.Split(gg[1], ";")
	for _, set := range sets {
		// After each time Elf shows the cubes, he puts them back.
		cubes := cubes(map[string]int{})

		colorCounts := strings.Split(set, ",")
		for _, cc := range colorCounts {
			ccp := str.MustSplit(strings.TrimSpace(cc), " ", 2)
			cubes[ccp[1]] = str.MustParseInt(ccp[0])
		}

		log.Debugf("checking game %d, cubes set: %v", id, cubes)

		// If a cube set doesn't fit as subset of available cubes, stop processing.
		if !cubes.subset(available) {
			return nil
		}
	}

	p.sum += id

	return nil
}

func (p *processor) result() (any, error) {
	return p.sum, nil
}
