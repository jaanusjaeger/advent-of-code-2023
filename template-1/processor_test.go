package main

import (
	"testing"

	"github.com/jaanusjaeger/advent-of-code-2023/io"
)

func TestProcessor(t *testing.T) {
	tests := []struct {
		input  string
		result any
	}{
		{
			input:  "example",
			result: -1, // TODO
		},
		{
			input:  "data",
			result: -1, // TODO
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
