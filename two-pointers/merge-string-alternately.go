package main

import "fmt"

func MergeAlternatelyDemo() {
	word1, word2 := "abc", "pqr"
	result := mergeAlternately(word1, word2)
	fmt.Println(result)
}

func mergeAlternately(word1 string, word2 string) string {
	l1, l2 := len(word1), len(word2)
	var result []byte
	i, j := 0, 0

	for i < l1 && j < l2 {
		result = append(result, word1[i])
		result = append(result, word2[j])
		i++
		j++
	}

	if i < l1 {
		result = append(result, word1[i:]...)
	}

	if j < l2 {
		result = append(result, word2[j:]...)
	}

	return string(result)
}
