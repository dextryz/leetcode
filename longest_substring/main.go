package main

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func lengthOfLongestSubstring(s string) int {

	m := make(map[rune]int)

    if len(s) == 1 {
        return 1
    }

	r := []rune(s)

	res, i, j := 0, 0, 0

	for j < len(r) {
		m[r[j]]++
        for m[r[j]] > 1 {
            m[r[i]]--
			i++
		}
        res = max(j-i+1, res)
		j++
	}

	return res
}
