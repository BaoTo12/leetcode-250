package main

import "fmt"

func BoyerMooreVotingAlgorithmDemo() {
	nums := []int{1, 2, 1, 3, 1, 1, 5} // 7 elements --> n / 2 = 3
	result := BoyerMooreVotingAlgorithm(nums)
	fmt.Println(result)
}

func BoyerMooreVotingAlgorithm(nums []int) int {
	votes := 0
	candidate := -1

	for _, v := range nums {
		if votes == 0 {
			candidate = v
			votes = 1
		} else {
			if candidate == v {
				votes += 1
			} else {
				votes -= 1
			}
		}
	}

	count := 0
	for _, v := range nums {
		if v == candidate {
			count++
		}
	}
	if count > len(nums)/2 {
		return candidate
	}

	return -1
}
