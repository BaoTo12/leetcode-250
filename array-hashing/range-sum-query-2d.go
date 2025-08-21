package main

import "fmt"

func SumRegionDemo() {
	matrix := [][]int{
		{3, 0, 1, 4, 2},
		{5, 6, 3, 2, 1},
		{1, 2, 0, 1, 5},
		{4, 1, 0, 1, 7},
		{1, 0, 3, 0, 5},
	}
	result := SumRegion(matrix, 2, 1, 4, 3)
	fmt.Println(result)
}

func PrefixSum(matrix [][]int) [][]int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return [][]int{}
	}
	m := len(matrix)
	n := len(matrix[0])
	prefixSum := make([][]int, m+1)
	for i := range prefixSum {
		prefixSum[i] = make([]int, n+1)
	}
	fmt.Println(prefixSum)
	for i := range m {
		for j := range n {
			prefixSum[i+1][j+1] =
				prefixSum[i+1][j] +
					prefixSum[i][j+1] -
					prefixSum[i][j] +
					matrix[i][j]
		}
	}
	return prefixSum
}

func SumRegion(matrix [][]int, row1 int, col1 int, row2 int, col2 int) int {
	prefixSum := PrefixSum(matrix)
	sum := prefixSum[row2+1][col2+1] - prefixSum[row2+1][col1] - prefixSum[row1][col2+1] + prefixSum[row1][col1]
	return sum
}
