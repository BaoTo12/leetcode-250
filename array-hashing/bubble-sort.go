package main

import "fmt"

func BubbleSortDemo() {
	nums := []int{5, 2, 3, 1}
	BubbleSort(nums)
	fmt.Println(nums)
}

func BubbleSort(nums []int) {
	n := len(nums)
	if n < 2 {
		return
	}
	for i := range nums {
		swap := false
		for j := range nums[:len(nums)-1-i] {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				swap = true
			}
		}
		if !swap {
			return
		}
	}
}
