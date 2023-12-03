package str

import (
	"fmt"
	"strconv"
	"strings"
)

// MustSplit splits the given string s by separator sep and panics if that there
// not are exactly n parts.
func MustSplit(s, sep string, n int) []string {
	res := strings.Split(s, sep)
	if len(res) != n {
		panic(fmt.Errorf("splitting %q by %q, expected %d parts, got %d", s, sep, n, len(res)))
	}
	return res
}

// MustParseInt parses string s to integer or panics if it can't parse.
func MustParseInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(fmt.Errorf("parsing %q to int: %w", s, err))
	}
	return i
}

func IsDigit(b byte) bool {
	return b >= '0' && b <= '9'
}
