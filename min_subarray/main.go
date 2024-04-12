package main

import "fmt"

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func minSubArrayLen(target int, nums []int) int {

    res := len(nums)

	var i, j, sum int

    for _, v := range nums {
        sum += v
    }
    if sum < target {
        return 0
    }

    sum = 0

	for j < len(nums) {

		sum += nums[j]

        fmt.Println(sum)

		for sum >= target {
			res = min(res, j-i+1)
			sum -= nums[i]
			i++
		}

		j++
	}

	return res
}
