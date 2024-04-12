package main

import "testing"

func Test(t *testing.T) {
	testCases := []struct {
		desc  string
		input string
		want  int
	}{
		{
			desc:  "A",
			input: "abcabcbb",
			want:  3,
		},
		{
			desc:  "B",
			input: "bbbbb",
			want:  1,
		},
		{
			desc:  "C",
			input: "pwwkew",
			want:  3,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := lengthOfLongestSubstring(tC.input)
			if tC.want != got {
				t.Errorf("want: %d, got: %d", tC.want, got)
			}
		})
	}
}
