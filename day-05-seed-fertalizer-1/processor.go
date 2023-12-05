package main

import (
	"github.com/jaanusjaeger/advent-of-code-2023/log"
	"github.com/jaanusjaeger/advent-of-code-2023/str"
)

type processor struct {
	seeds  []int
	next   bool
	groups []group
}

type group struct {
	name     string
	mappings []mapping
}

func (g group) transform(i int) int {
	for _, m := range g.mappings {
		t, ok := m.transform(i)
		if !ok {
			continue
		}
		return t
	}
	return i
}

type mapping struct {
	src   int
	dst   int
	count int
}

func (m mapping) transform(i int) (int, bool) {
	if i < m.src || i >= m.src+m.count {
		return -1, false
	}
	return m.dst + (i - m.src), true
}

func newProcessor() *processor {
	return &processor{}
}

func (p *processor) process(line string) error {
	if p.seeds == nil {
		seedsRow := str.MustSplit(line, ":", 2)
		seedsRaw := str.Split(seedsRow[1], " ")
		for _, s := range seedsRaw {
			p.seeds = append(p.seeds, str.MustParseInt(s))
		}
		return nil
	}

	if line == "" {
		p.next = true
		return nil
	}

	if p.next {
		groupRow := str.Split(line, " ")
		p.groups = append(p.groups, group{name: groupRow[0]})
		p.next = false
		return nil
	}

	group := &p.groups[len(p.groups)-1]
	mappingRow := str.MustSplit(line, " ", 3)
	group.mappings = append(group.mappings, mapping{
		dst:   str.MustParseInt(mappingRow[0]),
		src:   str.MustParseInt(mappingRow[1]),
		count: str.MustParseInt(mappingRow[2]),
	})

	return nil
}

func (p *processor) result() (any, error) {
	min := -1

	log.Debugf("GROUPS: %v", p.groups)

	for _, seed := range p.seeds {
		log.Debugf("seed %d", seed)
		val := seed
		for _, g := range p.groups {
			val = g.transform(val)
			log.Debugf("    %d by %q", val, g.name)
		}
		if min < 0 || val < min {
			min = val
		}
	}

	return min, nil
}
