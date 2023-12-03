package main

import (
	"testing"

	"github.com/jaanusjaeger/advent-of-code-2023/io"
	"github.com/jaanusjaeger/advent-of-code-2023/log"
)

func TestProcessor(t *testing.T) {
	log.SetDebug(false)

	tests := []struct {
		input  string
		result any
	}{
		{
			input:  "example",
			result: 2286,
		},
		{
			input:  "data",
			result: 55593,
		},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			p := newProcessor()

			err := io.LineByLine("input/"+test.input, p.process)
			if err != nil {
				t.Fatal(err)
			}

			result, err := p.result()
			if err != nil {
				t.Fatal(err)
			}
			if result != test.result {
				t.Errorf("want %v, got %v", test.result, result)
			}
		})
	}
}
