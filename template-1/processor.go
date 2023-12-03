package main

type processor struct {
}

func newProcessor() *processor {
	return &processor{}
}

func (p *processor) process(line string) error {
	// TODO
	return nil
}

func (p *processor) result() (any, error) {
	// TODO
	return nil, nil
}
