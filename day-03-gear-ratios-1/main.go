package main

import (
	"os"

	"github.com/jaanusjaeger/advent-of-code-2023/io"
	"github.com/jaanusjaeger/advent-of-code-2023/log"
)

func main() {
	input := io.MustGetArg()

	log.Infof("INPUT : %s", input)

	p := newProcessor()

	err := io.LineByLine(input, p.process)
	if err != nil {
		log.Errorf("reading input %q: %s", input, err)
		os.Exit(1)
	}

	result, err := p.result()
	if err != nil {
		log.Errorf("getting result: %s", err)
		os.Exit(1)
	}
	log.Infof("RESULT: %v", result)
}
