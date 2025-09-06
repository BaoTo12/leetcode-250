package main

import (
	"fmt"
	"strconv"
)

func CalPointsDemo() {
	operations := []string{"5", "-2", "4", "C", "D", "9", "+", "+"}
	fmt.Println(CalPoints(operations))
}

func CalPoints(operations []string) int {
	result := []int{}
	for _, str := range operations {
		switch str {
		case "+":
			previousSum := 0
			for _, v := range result[len(result)-2:] {
				previousSum += v
			}
			result = append(result, previousSum)
		case "C":
			result = result[:len(result)-1]
		case "D":
			result = append(result, result[len(result)-1]*2)
		default:
			integer, _ := strconv.Atoi(str)
			result = append(result, integer)
		}
	}
	sum := 0
	fmt.Println(result)
	for _, v := range result {
		sum += v
	}
	return sum
}
