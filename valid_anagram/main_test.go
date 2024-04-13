package main

import "testing"

func Test(t *testing.T) {
	testCases := []struct {
		desc string
		s    string
		t    string
		want bool
	}{
		{
			desc: "A",
			s:    "anagram",
			t:    "nagaram",
			want: true,
		},
		{
			desc: "B",
			s:    "rat",
			t:    "car",
			want: false,
		},
		{
			desc: "C",
			s:    "aacc",
			t:    "ccac",
			want: false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := isAnagram(tC.s, tC.t)
			if got != tC.want {
				t.Errorf("want: %v, got: %v", tC.want, got)
			}
		})
	}
}
