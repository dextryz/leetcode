package main

import (
	"reflect"
	"sort"
	"strings"
	"testing"
)

func sortAnagramSlices(slices [][]string) [][]string {
	for i := range slices {
		sort.Strings(slices[i])
	}
	sort.Slice(slices, func(i, j int) bool {
		return strings.Join(slices[i], ",") < strings.Join(slices[j], ",")
	})
	return slices
}

func Test(t *testing.T) {
	testCases := []struct {
		desc string
		strs []string
		want [][]string
	}{
		{
			desc: "A",
			strs: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			want: [][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}},
		},
		{
			desc: "B",
			strs: []string{""},
			want: [][]string{{""}},
		},
		{
			desc: "C",
			strs: []string{"a"},
			want: [][]string{{"a"}},
		},
		{
			desc: "D",
			strs: []string{"cab", "tin", "pew", "duh", "may", "ill", "buy", "bar", "max", "doc"},
			want: [][]string{{"max"}, {"buy"}, {"doc"}, {"may"}, {"ill"}, {"duh"}, {"tin"}, {"bar"}, {"pew"}, {"cab"}},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := groupAnagrams(tC.strs)
			if !reflect.DeepEqual(sortAnagramSlices(got), sortAnagramSlices(tC.want)) {
				t.Errorf("Failed %s: expected %v, got %v", tC.desc, tC.want, got)
			}
		})
	}
}
