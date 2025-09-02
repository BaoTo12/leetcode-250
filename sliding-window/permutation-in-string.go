package main

import (
	"fmt"
	"slices"
)

func CheckInclusionDemo() {
	s1, s2 := "adc", "dcda"
	fmt.Println(CheckInclusionV2(s1, s2))
}

func CheckInclusionV1(s1 string, s2 string) bool {
	sortedS1 := []byte(s1)
	slices.Sort(sortedS1)
	s1 = string(sortedS1)
	left := 0
	right := len(s1)
	for right <= len(s2) {
		str := []byte(string(s2[left:right]))
		slices.Sort(str)
		sortedStr := string(str)
		if sortedStr == s1 {
			return true
		}
		left++
		right++
	}

	return false
}

func CheckInclusionV2(s1 string, s2 string) bool {

	return false
}
