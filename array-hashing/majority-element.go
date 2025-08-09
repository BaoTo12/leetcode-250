package main

import "fmt"

func MajorityElementDemo() {
	// nums := []int{3, 2, 3}
	nums := []int{2, 2, 1, 1, 1, 2, 2}
	result := MajorityElement(nums)

	fmt.Println(result)
}

// ! Ver1
func MajorityElement(nums []int) int {
	ele := make(map[int]int)
	for _, v := range nums {
		ele[v] += 1
		if ele[v] > len(nums)/2 {
			return v
		}
	}
	return 0
}
