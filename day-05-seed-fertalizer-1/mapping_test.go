package main

import "testing"

func TestGroupTransform(t *testing.T) {
	g := group{
		name: "basic",
		mappings: []mapping{
			{
				src:   98,
				dst:   50,
				count: 2,
			},
			{
				src:   50,
				dst:   52,
				count: 48,
			},
		},
	}
	tests := []struct {
		name string
		in   int
		want int
	}{
		{
			name: "in-14",
			in:   14,
			want: 14,
		},
		{
			name: "in-50",
			in:   50,
			want: 52,
		},
		{
			name: "in-55",
			in:   55,
			want: 57,
		},
		{
			name: "in-79",
			in:   79,
			want: 81,
		},
		{
			name: "in-97",
			in:   97,
			want: 99,
		},
		{
			name: "in-98",
			in:   98,
			want: 50,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := g.transform(test.in)
			if got != test.want {
				t.Errorf("want %d, got %d", test.want, got)
			}
		})
	}
}
