package main

import (
	"github.com/jaanusjaeger/advent-of-code-2023/log"
	"github.com/jaanusjaeger/advent-of-code-2023/str"
)

type processor struct {
	times     []int
	distances []int
}

func newProcessor() *processor {
	return &processor{}
}

func (p *processor) process(line string) error {
	row := str.MustSplit(line, ":", 2)
	values := str.MustSplitInt(row[1], " ")

	switch row[0] {
	case "Time":
		p.times = values
	case "Distance":
		p.distances = values
	default:
		panic("unknown row type: " + row[0])
	}

	return nil
}

func (p *processor) result() (any, error) {
	var prod int = 1

	for i, time := range p.times {
		var count int
		targetDistance := p.distances[i]

		log.Debugf("time:     %d", time)
		log.Debugf("distance: %d", targetDistance)

		for speed := 0; speed < time; speed++ {
			distance := (time - speed) * speed

			log.Debugf("  speed:    %d", speed)
			log.Debugf("  distance: %d", distance)
			log.Debugf("    better?: %t", distance > targetDistance)

			if distance > targetDistance {
				count++
				continue
			}
		}
		log.Debugf("  count: %d", count)
		prod *= int(count)
	}

	return prod, nil
}
