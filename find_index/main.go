package main

// Getting the indexes correct was actually tricky.
func strStr(haystack string, needle string) int {

    // Encoding the length difference into the loop prevents an 
    // explicit bound check in the if-statement. Neet trick.
	for i := 0; i < len(haystack)-len(needle)+1; i++ {
		if needle == haystack[i:i+len(needle)] {
			return i
		}
	}

	return -1
}
