package main

import (
	"fmt"
)

func CharacterReplacementDemo() {
	s := "ABAB"
	k := 2
	fmt.Println(CharacterReplacement(s, k))
}

func CharacterReplacement(s string, k int) int {
	max_count := 0
	left := 0
	counter := make([]int, 26)
	res := 0

	for right, v := range s {
		idx := v - 'A'
		counter[idx]++
		if counter[idx] > max_count {
			max_count = counter[idx]
		}
		window := right - left + 1
		if window-max_count > k {
			counter[s[left]-'A'] -= 1
			left += 1
		}
		if (right - left + 1) > res {
			res = (right - left + 1)
		}
	}
	return res
}
