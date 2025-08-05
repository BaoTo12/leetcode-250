package main

import "fmt"

func ContainsDuplicateDemo() {
	s := []int{1, 2, 3, 1}
	fmt.Println(containsDuplicate(s))
}

func containsDuplicate(nums []int) bool {
	for _, v1 := range nums {
		for _, v2 := range nums {
			if v1 == v2 {
				return true
			}
		}
	}
	return false
}
