package main

import "fmt"

func SortColorsDemo() {
	nums := []int{2, 0, 2, 1, 1, 0}
	SortColors(nums)
	fmt.Println(nums)
}

func SortColors(nums []int) {
	// initialize counting array
	cntArr := make([]int, 2+1)
	for _, v := range nums {
		cntArr[v]++
	}
	// Cumulate the frequency of numbers
	for i := 1; i < len(cntArr); i++ {
		cntArr[i] += cntArr[i-1]
	}
	fmt.Println(cntArr)
	// assign output
	arrIndex := 0
	for i, v := range cntArr {
		for arrIndex < v {
			nums[arrIndex] = i
			arrIndex++
		}
	}

}
