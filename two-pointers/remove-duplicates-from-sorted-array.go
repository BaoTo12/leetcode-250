package main

import "fmt"

func RemoveDuplicatesDemo() {
	nums := []int{1, 1, 2}
	result := RemoveDuplicatesV1(nums)
	fmt.Println(result)
}

func RemoveDuplicatesV1(nums []int) int {
	i, j := 0, 0
	for range len(nums) {
		for j < len(nums) && nums[i] == nums[j] {
			j++
		}
		if j < len(nums) {
			i++
			nums[i] = nums[j]
		}
	}
	fmt.Println(nums)
	return i + 1
}
