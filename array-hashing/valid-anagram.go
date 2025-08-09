package main

import "fmt"

func isAnagramDemo() {
	// s := string("anagram")
	// t := string("nagaram")

	s := "rat"
	t := "car"
	fmt.Println(IsAnagram(s, t))
}

func IsAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	s1 := make(map[int32]int, len(s))
	for _, v := range s {
		s1[v] += 1

	}
	t1 := make(map[int32]int, len(t))
	for _, v := range t {
		t1[v] += 1
	}

	for _, v := range s {
		if s1[v] != t1[v] {
			return false
		}
	}
	return true
}
