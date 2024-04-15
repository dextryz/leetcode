package main

import (
	"strings"
)

func simplifyPath(path string) string {

    stack := []string{"/"}

    for _, v := range strings.Split(path, "/") {

        if v == "." || v == "" {
            continue
        } else if v == ".." {
            e := stack[len(stack)-1]
            if e != "/" {
                stack = stack[:len(stack)-1]
            }
        } else {
            stack = append(stack, v)
        }
    }

    return "/" + strings.Join(stack[1:], "/")
}
