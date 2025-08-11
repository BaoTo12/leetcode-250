package main

import (
	"fmt"
	"sort"
	"strings"
)

// Input: strs = ["eat","tea","tan","ate","nat","bat"]
// Output: [["bat"],["nat","tan"],["ate","eat","tea"]]
func GroupAnagramsDemo() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	result := groupAnagrams(strs)
	fmt.Println(result)
}

func groupAnagrams(strs []string) [][]string {
	if len(strs) == 0 || len(strs) == 1 {
		return [][]string{strs[:]}
	}
	groupAnagrams := map[string][]string{}
	// TODO Step1: Loop through strs
	for _, s := range strs {
		// TODO Step2: For each string we will sort the string to get the sorted key and put them into a slice
		// get sorted key
		sliceOfStr := strings.Split(s, "")
		sort.StringSlice(sliceOfStr).Sort()
		sortedKey := strings.Join(sliceOfStr, "")

		// add sorted key to groupAnagrams
		if _, err := groupAnagrams[sortedKey]; !err {
			groupAnagrams[sortedKey] = []string{}
		}
		groupAnagrams[sortedKey] = append(groupAnagrams[sortedKey], s)
	}
	fmt.Println(groupAnagrams)
	// TODO Step3: Return the slice of group anagram
	result := [][]string{}
	for _, ana := range groupAnagrams {
		result = append(result, ana)
	}
	return result
}
