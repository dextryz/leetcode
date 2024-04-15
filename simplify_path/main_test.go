package main

import "testing"

func Test(t *testing.T) {
	testCases := []struct {
		desc string
		path string
		want string
	}{
		{
			desc: "A",
			path: "/home/",
			want: "/home",
		},
		{
			desc: "B",
			path: "/../",
			want: "/",
		},
		{
			desc: "C",
			path: "/home//foo/",
			want: "/home/foo",
		},
		{
			desc: "D",
			path: "/a/./b/../../c/",
			want: "/c",
		},
		{
			desc: "E",
			path: "/a/../../b/../c//.//",
			want: "/c",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := simplifyPath(tC.path)
			if got != tC.want {
				t.Errorf("want: %v, got: %v", tC.want, got)
			}
		})
	}
}
