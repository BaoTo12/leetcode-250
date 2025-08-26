package main

import "fmt"

func mergeDemo() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	// nums1 := []int{0}
	m := 3
	num2 := []int{2, 5, 6}
	// num2 := []int{1}
	n := 3
	merge(nums1, m, num2, n)
	fmt.Println(nums1)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	i, j, length := m-1, n-1, m+n-1
	for i >= 0 && j >= 0 {
		if nums1[i] > nums2[j] {
			nums1[length] = nums1[i]
			i--
		} else {
			nums1[length] = nums2[j]
			j--
		}
		length--
	}
	for i >= 0 {
		nums1[length] = nums1[i]
		length--
		i--
	}
	for j >= 0 {
		nums1[length] = nums2[j]
		length--
		j--
	}
}
