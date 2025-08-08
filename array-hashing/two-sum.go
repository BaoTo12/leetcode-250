package main

import "fmt"

func TwoSumDemo() {
	nums := []int{2, 7, 11, 15}
	var target int = 9
	result := twoSum2(nums, target)

	if result != nil {
		fmt.Printf("First index: %d", result[0])
		fmt.Println()
		fmt.Printf("Second index: %d", result[1])
	}
}

// ! Ver 1
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

// ! Ver2
// {2, 7, 11, 15}
func twoSum2(nums []int, target int) []int {
	pair_idx := make(map[int]int)
	for i, v := range nums {
		complement := target - v
		if complementIndex, ok := pair_idx[complement]; ok {
			return []int{complementIndex, i}
		}
		pair_idx[v] = i
	}
	return nil
}
