package main

import "fmt"

func RotateArrayDemo() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 4
	// RotateV1(nums, k)
	// RotateV2(nums, k)
	RotateV3(nums, k)
	fmt.Println(nums)
}

func RotateV1(nums []int, k int) {
	n := len(nums)
	k = k % n
	rotatedArray := make([]int, n)
	for i, v := range nums {
		index := (i + k) % n
		rotatedArray[index] = v
	}
	for i := range nums {
		nums[i] = rotatedArray[i]
	}
	fmt.Println(nums)
}

func RotateV2(nums []int, k int) {
	n := len(nums)
	k = k % n
	k = k % n
	if k == 0 {
		return
	}

	rotated := append(nums[n-k:], nums[:n-k]...)

	copy(nums, rotated)
}

func RotateV3(nums []int, k int) {
	n := len(nums)
	k = k % n
	// reverse the whole array
	reverse(nums, 0, n-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, n-1)
}

func reverse(nums []int, left int, right int) {
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}
