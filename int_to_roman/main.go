package main

import (
	"fmt"
	"strconv"
)

func intToRoman(num int) string {

	lookup := map[int]string{
		1:    "I",
		4:    "IV",
		5:    "V",
		9:    "IX",
		10:   "X",
		40:   "XL",
		50:   "L",
		90:   "XC",
		100:  "C",
		400:  "CD",
		500:  "D",
		900:  "CM",
		1000: "M",
	}

	fmt.Println(lookup)

	stack := []int{}

	s := strconv.Itoa(num)

	r := []rune(s)

	d := 1

	for i := len(r) - 1; i >= 0; i-- {

		v := d * int(r[i]-'0')

		stack = append(stack, v)

		fmt.Println(v)

		d *= 10
	}

    res := ""
    for i := len(stack)-1; i >= 0; i-- {
        fmt.Printf("%d ", stack[i])

        v, ok := lookup[stack[i]]
        if ok {
            res += v
        }

    }
    fmt.Println("")

	return res
}
