package main

import "fmt"

func ConcatenationOfArray() {
	s1 := []int{1, 2, 1}
	s2 := getConcatenation(s1)
	fmt.Println(s2)
}

func getConcatenation(nums []int) []int {
	s2 := make([]int, len(nums)*2)
	for i, v := range nums {
		s2[i] = v
		s2[i+len(nums)] = v
	}
	return s2
}
