package main

import "fmt"

func twoSumDemo() {
	// numbers := []int{2, 7, 11, 15}
	numbers := []int{2, 3, 4}
	target := 6
	result := twoSum(numbers, target)
	fmt.Println(result)
}

func twoSum(numbers []int, target int) []int {
	arr := map[int]int{}
	for i, v := range numbers {
		need := target - v
		if value, err := arr[need]; err {
			return []int{value + 1, i + 1}
		}
		arr[v] = i
	}
	return nil
}
