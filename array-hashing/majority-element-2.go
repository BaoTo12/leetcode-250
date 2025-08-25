package main

import (
	"fmt"
	"math"
)

func MajorityElementIIDemo() {
	// nums := []int{3, 2, 3}
	// nums := []int{1}
	nums := []int{1, 2}
	result := MajorityElementII(nums)
	fmt.Println(result)
}

// TODO Simple Version
// func MajorityElementII(nums []int) []int {
// 	threshold := len(nums) / 3
// 	counter := map[int]int{}
// 	for _, v := range nums {
// 		counter[v]++
// 	}
// 	result := []int{}
// 	for key, val := range counter {
// 		if val > threshold {
// 			result = append(result, key)
// 		}
// 	}
// 	return result
// }

// TODO Boyer-Moore Majority Voting Algorithm
func MajorityElementII(nums []int) []int {
	ele1 := math.MaxInt
	ele2 := math.MaxInt

	freq1 := 0
	freq2 := 0

	for _, val := range nums {
		if freq1 == 0 && val != ele2 {
			ele1 = val
			freq1++
		} else if freq2 == 0 && val != ele1 {
			ele2 = val
			freq2++
		} else if val == ele1 {
			freq1++
		} else if val == ele2 {
			freq2++
		} else {
			freq1--
			freq2--
		}
	}

	freq1 = 0
	freq2 = 0
	for _, val := range nums {
		if val == ele1 {
			freq1++
		} else if val == ele2 {
			freq2++
		}
	}

	var ans []int
	n := len(nums)
	if freq1 > n/3 {
		ans = append(ans, ele1)
	}
	if freq2 > n/3 {
		ans = append(ans, ele2)
	}

	return ans
}

// ! Explain
/*
! Kết nối với thuật toán Boyer–Moore mở rộng
Thuật toán mô phỏng quá trình huỷ nhóm một cách hiệu quả trong một lần duyệt:

	Duy trì k-1 ứng viên và k-1 bộ đếm.

	Khi gặp một phần tử bằng ứng viên hiện có → tăng bộ đếm tương ứng.

	Khi gặp phần tử khác và có bộ đếm 0 → gán làm ứng viên mới.

	Khi gặp phần tử khác và tất cả k-1 bộ đếm đều > 0 → giảm tất cả k-1 bộ đếm (tương đương với huỷ một nhóm k phần tử khác nhau, nhưng thuật toán dùng k-1 vì nhóm thực tế có kích thước k gồm phần tử hiện tại + k-1 ứng viên khác).

Sau khi chạy xong, những ứng viên còn đếm dương là các loại có khả năng xuất hiện > n/k. Vì huỷ nhóm đã loại bỏ phần lớn các loại hiếm, chỉ còn ≤ k-1 loại có thể vượt ngưỡng — nên thuật toán chỉ cần thử tối đa k-1 ứng viên trong lần kiểm tra cuối cùng (đếm thực tế).
! Chứng minh ngắn cho tính đúng đắn (idea)
	Không thể có ≥ k giá trị, mỗi cái > n/k (bất đẳng thức chứng minh mâu thuẫn).
	Thuật toán huỷ nhóm sẽ không vô tình loại mất một giá trị có count > n/k: mỗi lần huỷ giảm count của mọi loại trong nhóm cùng 1 đơn vị, nên giá trị thực sự lớn hơn n/k vẫn còn tồn dư sau khi huỷ hết các nhóm có thể. Vì phần dư chỉ có ≤ k-1 loại, mọi loại > n/k phải nằm trong những loại đó.
	Do đó chỉ cần kiểm tra các ứng viên cuối cùng bằng một lần đếm thực tế để lọc ra những giá trị thật sự > n/k.

*/
