package main

import "fmt"

func LengthOfLongestSubstringDemo() {
	s := "abcabcbb"
	fmt.Println(LengthOfLongestSubstring(s))
}

func LengthOfLongestSubstring(s string) int {
	lastSeen := make(map[byte]int)
	start := 0
	maxLength := 0

	for i, char := range s {
		if lastIndex, ok := lastSeen[byte(char)]; ok && lastIndex >= start {
			start = lastIndex + 1
		}
		lastSeen[byte(char)] = i
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
	}

	return maxLength
}
