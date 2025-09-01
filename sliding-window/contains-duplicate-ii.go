package main

import (
	"fmt"
)

func ContainsNearbyDuplicateDemo() {
	nums := []int{1, 2, 3, 1}
	k := 3
	fmt.Println(ContainsNearbyDuplicate(nums, k))
}

func ContainsNearbyDuplicate(nums []int, k int) bool {
	arrMap := make(map[int]int)
	for i, v := range nums {
		if j, ok := arrMap[v]; ok {
			if i-j <= k {
				return true
			}
		}
		arrMap[v] = i
	}
	return false
}
