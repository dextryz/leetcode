package main

import "testing"

func assert(x, y []int) bool {
	if len(x) != len(y) {
		return false
	}
	for i := 0; i < len(x); i++ {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}

func Test(t *testing.T) {

	testCases := []struct {
		desc string
		nums []int
		want []int
		k    int
	}{
		{
			desc: "",
			nums: []int{1, 2, 3, 4, 5, 6, 7},
			k:    3,
			want: []int{5, 6, 7, 1, 2, 3, 4},
		},
		{
			desc: "",
			nums: []int{-1, -100, 3, 99},
			k:    2,
			want: []int{3, 99, -1, -100},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			rotate(tC.nums, tC.k)
			if !assert(tC.nums, tC.want) {
				t.Errorf("want: %v, got: %v", tC.want, tC.nums)
			}
		})
	}
}
