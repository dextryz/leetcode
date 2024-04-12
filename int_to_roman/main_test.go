package main

import "testing"

func Test(t *testing.T) {
	testCases := []struct {
		desc  string
		want  string
		input int
	}{
		{
			desc:  "",
			input: 3,
			want:  "III",
		},
		{
			desc:  "",
			input: 58,
			want:  "LVIII",
		},
		{
			desc:  "",
			input: 1994,
			want:  "MCMXCIV",
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := intToRoman(tC.input)
			if got != tC.want {
				t.Errorf("want: %s, got: %s", tC.want, got)
			}
		})
	}
}
