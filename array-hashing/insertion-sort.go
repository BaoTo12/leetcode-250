package main

import "fmt"

func InsertionSortDemo() {
	nums := []int{12, 11, 13, 5, 6}
	InsertionSort(nums)
	fmt.Println(nums)
}

func InsertionSort(nums []int) {
	n := len(nums)
	if n < 2 {
		return
	}
	for i := 1; i < n; i++ {
		key := nums[i]
		j := i - 1
		for j >= 0 && nums[j] > key {
			nums[j+1] = nums[j]
			j--
		}
		nums[j+1] = key
	}
}
