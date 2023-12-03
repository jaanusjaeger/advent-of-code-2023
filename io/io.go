package io

import (
	"bufio"
	"fmt"
	"os"
)

func MustGetArg() string {
	if len(os.Args) < 2 {
		panic("must give cmd line argument")
	}
	return os.Args[1]
}

func LineByLine(file string, cb func(lile string) error) error {
	f, err := os.Open(file)
	if err != nil {
		return fmt.Errorf("opening file: %w", err)
	}

	scanner := bufio.NewScanner(f)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	// https://stackoverflow.com/questions/8757389/reading-a-file-line-by-line-in-go/16615559#16615559
	for scanner.Scan() {
		l := scanner.Text()
		err := cb(l)
		if err != nil {
			return fmt.Errorf("processing line %q: %w", l, err)
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("reading file: %w", err)
	}

	return nil
}
