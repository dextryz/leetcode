package main

import "testing"

func Test(t *testing.T) {
	testCases := []struct {
		desc string
		str1 string
		str2 string
		want string
	}{
		{
			desc: "A",
			str1: "ABCABC",
			str2: "ABC",
			want: "ABC",
		},
		{
			desc: "B",
			str1: "ABABAB",
			str2: "ABAB",
			want: "AB",
		},
		{
			desc: "C",
			str1: "LEET",
			str2: "CODE",
			want: "",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := gcdOfStrings(tC.str1, tC.str2)
			if got != tC.want {
				t.Errorf("want: %v, got: %v", tC.want, got)
			}
		})
	}
}
