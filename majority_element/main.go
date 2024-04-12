package main

import "fmt"

// https://leetcode.com/problems/majority-element/description/?envType=study-plan-v2&envId=top-interview-150

func majorityElement(nums []int) int {

	m := make(map[int]int)

	for _, v := range nums {
		m[v]++
	}

	for k, v := range m {
		if v > len(nums)/2 {
			return k
		}
	}

	return -1
}

func main() {

	r := majorityElement([]int{3, 2, 3})
	fmt.Println(r) // 3

	r = majorityElement([]int{2, 2, 1, 1, 1, 2, 2})
	fmt.Println(r) // 2
}
