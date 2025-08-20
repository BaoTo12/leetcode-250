package main

import "fmt"

func TopKFrequentDemo() {
	nums := []int{1, 1, 1, 2, 2, 3}
	k := 2
	result := TopKFrequentBucket(nums, k)
	fmt.Println(result)
}

func TopKFrequentBucket(nums []int, k int) []int {
	counter := map[int]int{}

	for _, v := range nums {
		counter[v]++
	}
	fmt.Println(counter)
	buckets := make([][]int, len(nums)+1)
	for val, key := range counter {
		buckets[key] = append(buckets[key], val)
	}
	res := []int{}
	for freq := len(nums); freq >= 0 && len(res) < k; freq-- {
		for _, num := range buckets[freq] {
			res = append(res, num)
			if len(res) == k {
				return res
			}
		}
	}
	fmt.Println(buckets)
	return nil
}
