package main

import "testing"

func Test(t *testing.T) {
	testCases := []struct {
		desc  string
		input string
		want  int
	}{
		{
			desc:  "D",
			input: "IV",
			want:  4,
		},
		{
			desc:  "A",
			input: "III",
			want:  3,
		},
		{
			desc:  "B",
			input: "LVIII",
			want:  58,
		},
		{
			desc:  "C",
			input: "MCMXCIV",
			want:  1994,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
            got := romanToInt(tC.input)
            if tC.want != got {
                t.Errorf("want: %d, got %d", tC.want, got)
            }
		})
	}
}
