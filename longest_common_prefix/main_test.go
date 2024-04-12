package main

import "testing"

func Test(t *testing.T) {

	testCases := []struct {
		desc string
		want string
		strs []string
	}{
		{
			desc: "A",
			strs: []string{"flower", "flow", "flight"},
			want: "fl",
		},
		{
			desc: "B",
			strs: []string{"dog", "racecar", "car"},
			want: "",
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := longestCommonPrefix(tC.strs)
			if got != tC.want {
				t.Errorf("want: %s, got: %s", tC.want, got)
			}
		})
	}
}
