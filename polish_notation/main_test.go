package main

import "testing"

func Test(t *testing.T) {
	testCases := []struct {
		desc   string
		tokens []string
		want   int
	}{
		{
			desc:   "A",
			tokens: []string{"2", "1", "+", "3", "*"},
			want:   9,
		},
		{
			desc:   "B",
			tokens: []string{"4", "13", "5", "/", "+"},
			want:   6,
		},
		{
			desc:   "C",
			tokens: []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"},
			want:   22,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := evalRPN(tC.tokens)
			if got != tC.want {
				t.Errorf("want: %v, got: %v", tC.want, got)
			}
		})
	}
}
