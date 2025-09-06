package main

import (
	"fmt"
	"math"
)

func MinSubArrayLenDemo() {
	nums := []int{1, 1, 1, 1, 1, 1, 1, 1}
	target := 11
	fmt.Println(MinSubArrayLen(target, nums))
}

func MinSubArrayLen(target int, nums []int) int {
	left := 0
	cur_len := math.MaxInt
	sum := 0

	for right, value := range nums {
		sum += value

		for sum >= target {
			if right-left+1 < cur_len {
				cur_len = right - left + 1
			}
			sum -= nums[left]
			left += 1
		}
	}
	if cur_len == math.MaxInt {
		return 0
	} else {
		return cur_len
	}
}
