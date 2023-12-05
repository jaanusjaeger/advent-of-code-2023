package main

import (
	"github.com/jaanusjaeger/advent-of-code-2023/log"
	"github.com/jaanusjaeger/advent-of-code-2023/str"
)

type processor struct {
	seeds  []seedRange
	next   bool
	groups []group
}

type seedRange struct {
	from  int
	count int
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
		for i := 0; i < len(seedsRaw); i += 2 {
			p.seeds = append(p.seeds, seedRange{
				from:  str.MustParseInt(seedsRaw[i]),
				count: str.MustParseInt(seedsRaw[i+1]),
			})
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

	for _, sr := range p.seeds {
		log.Debugf("sr: %+v", sr)
		for i := 0; i < sr.count; i++ {
			val := sr.from + i
			for _, g := range p.groups {
				val = g.transform(val)
			}
			if min < 0 || val < min {
				min = val
			}
		}
	}

	return min, nil
}
