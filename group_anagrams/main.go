package main

// This is fucking genious
type Key [26]byte

func groupAnagrams(strs []string) [][]string {

	group := make(map[Key][]string)

	for _, word := range strs {
		var key Key
		for _, v := range word {
			key[v-'a']++
		}
		group[key] = append(group[key], word)
	}

	res := [][]string{}
	for _, v := range group {
		res = append(res, v)
	}

	return res
}
