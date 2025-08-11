package main

import "fmt"

func SelectionSortDemo() {
	nums := []int{64, 25, 12, 22, 11}
	SelectionSort(nums)
	fmt.Println(nums)
}

func SelectionSort(nums []int) {
	n := len(nums)
	if n < 2 {
		return
	}
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[minIdx] {
				minIdx = j
			}
		}
		// swap here
		if minIdx != i {
			nums[i], nums[minIdx] = nums[minIdx], nums[i]
		}
	}
}
