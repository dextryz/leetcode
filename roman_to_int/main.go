package main

// Using the multiple if-statements is confusing.
// Instead, look at how the pattern is encoded within the math.
func romanToInt(s string) int {

    lookup := map[rune]int{
        'I': 1,
        'V': 5,
        'X': 10,
        'L': 50,
        'C': 100,
        'D': 500,
        'M': 1000,
    }

    var v, cv, lv int

    for i := len(s)-1; i >= 0; i-- {
        cv = lookup[rune(s[i])]
        if cv < lv {
            v -= cv
        } else {
            v += cv
        }
        lv = cv
    }

    return v
}
