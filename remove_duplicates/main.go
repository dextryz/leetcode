package main

import "fmt"

func removeDuplicates(nums []int) int {

	i, j := 0, 1

	curr := nums[i]

	for j < len(nums) {

		if curr != nums[j] {
			i++
			nums[i] = nums[j]
			curr = nums[i]
		}

		j++
	}

	return i + 1
}

func main() {

	r := removeDuplicates([]int{1, 1, 2})
	fmt.Println(r) // 2, [1,2,_]

	r = removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4})
	fmt.Println(r) // 5, [0,1,2,3,4,_,_,_,_,_]
}
