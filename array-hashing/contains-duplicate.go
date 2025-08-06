package main

import (
	"fmt"
)

func ContainsDuplicateDemo() {
	s := []int{1, 2, 3, 1}
	// s1 := []int{1, 2, 3, 4}
	fmt.Println(CheckLengthToCheckDuplicate(s))
	// fmt.Println(CheckLengthToCheckDuplicate(s1))
}

func containsDuplicate(nums []int) bool {
	seen := make(map[int]struct{}, len(nums))
	for _, v := range nums {
		if _, existed := seen[v]; existed {
			return true
		}
		seen[v] = struct{}{}
	}
	return false
}
func CheckLengthToCheckDuplicate(nums []int) bool {
	seen := make(map[int]struct{}, len(nums))
	for _, v := range nums {
		seen[v] = struct{}{}
	}
	return len(seen) < len(nums)
}
