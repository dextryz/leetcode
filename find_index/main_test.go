package main

import "testing"

func Test(t *testing.T) {

	testCases := []struct {
		desc     string
		needle   string
		haystack string
		want     int
	}{
		{
			desc:     "A",
			needle:   "sad",
			haystack: "sadbutsad",
			want:     0,
		},
		{
			desc:     "B",
			needle:   "leeto",
			haystack: "leetcode",
			want:     -1,
		},
		{
			desc:     "C",
			needle:   "ll",
			haystack: "hello",
			want:     2,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := strStr(tC.haystack, tC.needle)
			if got != tC.want {
				t.Errorf("want: %d, got: %d", tC.want, got)
			}
		})
	}
}
