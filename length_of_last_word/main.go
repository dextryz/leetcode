package main

import (
	"fmt"
)

// 1. Loop form the back until the right pointer is at the first letter
// 2. Start left pointer at this index, then loop until first space
// 3. The word length the diff between the right and left pointers.
func lengthOfLastWord(s string) int {

	r := len(s) - 1

	for s[r] == ' ' {
		r--
	}

	l := r

    // TODO: This was actually pretty tricky.
	for l > 0 && s[l-1] != ' ' {
		l--
	}

	return r - l + 1

}

func main() {

	r := lengthOfLastWord("Hello World")
	fmt.Println(r) // 5

	r = lengthOfLastWord("   fly me   to   the moon  ")
	fmt.Println(r) // 4

}
