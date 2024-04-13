package main

import "testing"

func Test(t *testing.T) {
	testCases := []struct {
		desc    string
		s       string
		pattern string
		want    bool
	}{
		{
			desc:    "A",
			s:       "dog cat cat dog",
			pattern: "abba",
			want:    true,
		},
		{
			desc:    "B",
			s:       "dog cat cat fish",
			pattern: "abba",
			want:    false,
		},
		{
			desc:    "C",
			s:       "dog cat cat dog",
			pattern: "aaaa",
			want:    false,
		},
		{
			desc:    "D",
			s:       "dog dog dog dog",
			pattern: "abba",
			want:    false,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := wordPattern(tC.pattern, tC.s)
			if got != tC.want {
				t.Errorf("want: %v, got: %v", tC.want, got)
			}
		})
	}
}
