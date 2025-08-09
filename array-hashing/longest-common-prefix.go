package main

import (
	"fmt"
	"sort"
)

func LongestCommonPrefixDemo() {
	strs := []string{"flower", "flow", "flight"}
	result := LongestCommonPrefix(strs)
	fmt.Println(result)
}

func LongestCommonPrefix(strs []string) string {
	sort.Slice(strs, func(i, j int) bool {
		return len(strs[i]) < len(strs[j])
	})
	prefix := strs[0]
	prefix_length := len(prefix)

	for _, s := range strs[1:] {
		for prefix != s[0:prefix_length] {
			prefix_length -= 1
			if prefix_length == 0 {
				return ""
			}
			prefix = prefix[0:prefix_length]
		}
	}

	return prefix
}
