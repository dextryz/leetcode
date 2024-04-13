package main

import "strings"

func wordPattern(pattern string, s string) bool {

	m := make(map[rune]string)
	m1 := make(map[string]rune)
	words := strings.Fields(s)
    patterns := []rune(pattern)

    if len(patterns) != len(words) {
        return false
    }

	for i, v := range []rune(pattern) {
		if _, ok := m[v]; ok {
			if m[v] != words[i] {
				return false
			}
		}
        m[v] = words[i]
	}

    for i, v := range words {
        if _, ok := m1[v]; ok {
            if m1[v] != patterns[i] {
                return false
            }
        }
        m1[v] = patterns[i]
    }

	return true
}
