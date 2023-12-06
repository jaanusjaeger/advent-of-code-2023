package main

import (
	"strings"

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
	values := str.Split(row[1], " ")
	valueRaw := strings.Join(values, "")
	value := str.MustParseInt(valueRaw)

	switch row[0] {
	case "Time":
		p.times = append(p.times, value)
	case "Distance":
		p.distances = append(p.distances, value)
	default:
		panic("unknown row type: " + row[0])
	}

	return nil
}

func (p *processor) result() (any, error) {
	var prod int64 = 1

	var logCount int64
	debug := func(format string, args ...any) {
		if logCount%100000 == 0 {
			log.Debugf(format, args...)
		}
		logCount++
	}

	for i, time := range p.times {
		var count int
		targetDistance := p.distances[i]

		debug("time:     %d", time)
		debug("distance: %d", targetDistance)

		for speed := 0; speed < time; speed++ {
			distance := (time - speed) * speed

			debug("  speed:    %d", speed)
			debug("  distance: %d", distance)
			debug("    better?: %t", distance > targetDistance)

			if distance > targetDistance {
				count++
				continue
			}
		}
		debug("  count: %d", count)
		prod *= int64(count)
	}

	return prod, nil
}
