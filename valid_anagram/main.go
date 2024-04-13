package main

func isAnagram(s string, t string) bool {

    chars := make(map[rune]int)
    
    for _, v := range s {
        chars[v]++
    }
    
    for _, v := range t {
        if number, exists := chars[v]; !exists || number == 0 {
            return false
        }
        chars[v]--
    }
    
    return len(s) == len(t)
}
