package main

import (
	"fmt"
	"slices"
)

func ThreeSumDemo() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	result := ThreeSum(nums)
	fmt.Println(result)
}

func ThreeSum(nums []int) [][]int {
	slices.Sort(nums)
	result := [][]int{}
	for i := range len(nums) {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j := i + 1
		k := len(nums) - 1

		for j < k {
			total := nums[i] + nums[j] + nums[k]
			if total == 0 {
				result = append(result, []int{nums[i], nums[j], nums[k]})
				j += 1
				for nums[j] == nums[j-1] && j < k {
					j += 1
				}
			} else if total > 0 {
				k -= 1
			} else {
				j += 1
			}
		}
	}
	return result
}
