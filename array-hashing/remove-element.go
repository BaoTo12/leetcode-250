package main

import "fmt"

func RemoveElementDemo() {
	// nums := []int{3, 2, 2, 3}
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	// nums := []int{}
	// nums := []int{1}

	val := 2

	result := RemoveElement(nums, val)
	fmt.Println(result)
}

func RemoveElement(nums []int, val int) int {
	k := 0
	for i, v := range nums {
		if v != val {
			nums[k] = nums[i]
			k++
		}
	}
	fmt.Println(nums)
	return k
}
