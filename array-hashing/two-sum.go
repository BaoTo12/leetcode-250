package main

import "fmt"

func TwoSumDemo() {
	nums := []int{2, 7, 11, 15}
	var target int = 9
	result := twoSum(nums, target)

	if result != nil {
		fmt.Printf("First index: %d", result[0])
		fmt.Println()
		fmt.Printf("Second index: %d", result[1])
	}
}

func twoSum(nums []int, target int) []int {
	for i1, val1 := range nums {
		for i2, val2 := range nums {
			if val1+val2 == target {
				return []int{i1, i2}
			}
		}
	}
	return nil
}
