package main

import "fmt"

// TODO: This is Lomuto Partition' approach
// func QuickSortDemo() {
// 	nums := []int{1, 2, 3, 4, 5, 6, 7}
// 	// nums := []int{4, 3, 1, 2, 5, 9, 7, 10, 6}
// 	QuickSort(nums, 0, len(nums)-1)
// 	fmt.Println(nums)
// }

// func partition(nums []int, low int, high int) int {
// 	pivot := nums[high]

// 	i := low - 1
// 	for j := low; j < high; j++ {
// 		if nums[j] < pivot {
// 			i++
// 			nums[i], nums[j] = nums[j], nums[i]
// 		}
// 	}
// 	nums[i+1], nums[high] = nums[high], nums[i+1]
// 	return i + 1
// }

// func QuickSort(nums []int, low int, high int) {
// 	if low < high {
// 		pivotIndex := partition(nums, low, high)

// 		QuickSort(nums, low, pivotIndex-1)
// 		QuickSort(nums, pivotIndex+1, high)
// 	}
// }

// TODO: Enhanced Version

func QuickSortDemo() {
	nums := []int{4, 3, 1, 2, 5, 9, 7, 10, 6}
	fmt.Println(nums)
}

func medianOfThree(nums []int, low, high int) int {
	mid := low + (high-low)/2
	a, b, c := nums[low], nums[mid], nums[high]
	// find median value's index
	if b <= c {
		if c <= a {
			return high
		}
		if b >= a {
			return mid
		}
		return low
	} else { // a > b
		if c >= a {
			return high
		}
		if b >= a {
			return low
		}
		return mid
	}
}

func Partition(nums []int, low int, high int) int {
	return 0
}

func QuickSort(nums []int, low int, high int) {

}

// ! Choosing pivot
/*
	- Choose the first or last element
	- Pick a random element
	- Pick the median
*/

// ! Partition algorithm
/*
	- Native Partition: Here we create copy of the array. First put all smaller elements and then all greater. Finally we copy the temporary array back to original array. This requires O(n) extra space.
	- Lomuto Partition: We have used this partition in this article. This is a simple algorithm, we keep track of index of smaller elements and keep swapping. We have used it here in this article because of its simplicity.
	- Hoare's Partition: This is the fastest of all. Here we traverse array from both sides and keep swapping greater element on left with smaller on right while the array is not partitioned. Please refer Hoare’s vs Lomuto for details.

*/
