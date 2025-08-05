package main

import "fmt"

func ContainsDuplicateDemo() {
	// s := []int{1, 2, 3, 1}
	s1 := []int{1, 2, 3, 4}
	fmt.Println(containsDuplicate(s1))
}

func containsDuplicate(nums []int) bool {
	for i := range nums {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				return true
			}
		}
	}
	return false
}
