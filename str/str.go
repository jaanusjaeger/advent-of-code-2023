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

// Split is a special version of 'strings.Split' that removes empty values.
func Split(s, sep string) []string {
	return RmEmpty(strings.Split(s, sep))
}

// MustSplitInt calls 'Split' and then 'MustParseInt' on each resulting value.
func MustSplitInt(s, sep string) []int {
	vals := Split(s, sep)
	result := make([]int, len(vals))
	for i, v := range vals {
		result[i] = MustParseInt(v)
	}
	return result
}

func IsDigit(b byte) bool {
	return b >= '0' && b <= '9'
}

func RmEmpty[T comparable](vv []T) []T {
	var result []T
	var empty T

	for _, v := range vv {
		if v == empty {
			continue
		}
		result = append(result, v)
	}

	return result
}
