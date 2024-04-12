package main

// Without using the STD, the key is to leverage slices
// Look at the simplicity of the loop to move the pointer.
func longestCommonPrefix(strs []string) string {

	r := strs[0]

    for _, w := range strs {
        i := 0
        for i < len(w) && i < len(r) && w[i] == r[i] {
            i++
        }
        r = r[:i]
    }

	return r
}
