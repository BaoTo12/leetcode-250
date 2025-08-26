package main

import "fmt"

func IsPalindromeDemo() {
	s := "A man, a plan, a canal: Panama"
	fmt.Println(isPalindrome(s))
}

func isAlnum(b byte) bool {
	// '0'-'9' => 48-57, 'A'-'Z' => 65-90, 'a'-'z' => 97-122
	if b >= '0' && b <= '9' {
		return true
	}
	if b >= 'A' && b <= 'Z' {
		return true
	}
	if b >= 'a' && b <= 'z' {
		return true
	}
	return false
}
func toLower(b byte) byte {
	if b >= 'A' && b <= 'Z' {
		return b + ('a' - 'A')
	}
	return b
}

func isPalindrome(s string) bool {
	i := 0
	j := len(s) - 1
	for i < j {
		// move i forward to next alnum
		for i < j && !isAlnum(s[i]) {
			i++
		}
		// move j backward to prev alnum
		for i < j && !isAlnum(s[j]) {
			j--
		}
		if toLower(s[i]) != toLower(s[j]) {
			return false
		}
		i++
		j--
	}
	return true
}
