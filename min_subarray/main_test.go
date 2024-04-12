package main

import "testing"

func Test(t *testing.T) {

	testCases := []struct {
		desc   string
		nums   []int
		target int
		want   int
	}{
		{
			desc:   "A",
			nums:   []int{2, 3, 1, 2, 4, 3},
			target: 7,
			want:   2,
		},
		{
			desc:   "B",
			nums:   []int{1, 4, 4},
			target: 4,
			want:   1,
		},
		{
			desc:   "C",
			nums:   []int{1, 1, 1, 1, 1, 1, 1, 1},
			target: 11,
			want:   0,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := minSubArrayLen(tC.target, tC.nums)
			if got != tC.want {
				t.Errorf("want: %d, got: %d", tC.want, got)
			}
		})
	}
}
