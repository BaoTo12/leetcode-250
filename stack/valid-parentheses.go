package main

import "fmt"

func IsValidDemo() {
	s := "([])"
	fmt.Println(IsValid(s))
}

func IsValid(s string) bool {
	stack := []rune{}

	for _, c := range s {
		if c == '{' || c == '(' || c == '[' {
			stack = append(stack, c)
		} else {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if c == ')' && top != '(' {
				return false
			}
			if c == '}' && top != '{' {
				return false
			}
			if c == ']' && top != '[' {
				return false
			}
		}
	}
	return true
}
