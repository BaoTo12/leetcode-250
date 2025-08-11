package main

import "fmt"

func SortArrayDemo() {
	nums := []int{5, 2, 3, 1}
	result := sortArray(nums)
	fmt.Println(result)
}

func sortArray(nums []int) []int {

	return nil
}

// ! Comparison-based sorting (dựa trên so sánh)
// bubble sort: so sánh và đổi chỗ phần tử kề nhau cho đến khi mảng được sắp xếp. Rất trực quan, nhưng chậm (O(n²)), thường chỉ để học
// selection sort: ỗi lần chọn phần tử nhỏ nhất trong phần chưa sắp xếp, rồi swap vào vị trí thích hợp. O(n²), dùng khi write-cost cao vì số lượng “swap” tối thiểu
// insertion sort: kiểu như chơi bài: mỗi lần chèn phần tử vào vị trí đúng của phần đã sắp xếp. Thích hợp với mảng gần như đã sắp. O(n²) worst-case nhưng hiệu quả thực (adaptive)
// merge sort:  chia để trị, chia mảng thành hai phần, sort rồi merge lại. O(n log n) time, cần extra space O(n), nhưng là stable sort
// quick sort: dùng pivot để phân vùng, rồi đệ quy sort từng phần. Thường rất nhanh, O(n log n) trung bình, O(n²) worst-case, không stable
// heap sort:  xây heap rồi dần dần trích phần tử lớn nhất—O(n log n) best/average/worst, in-place, nhưng không stable

// ! Non-comparison sorting (không dùng so sánh)
// Counting Sort: dùng mảng đếm để xác định vị trí của phần tử. Linear O(n + k) với k = range of keys. Không stable nếu implement in-place, stable nếu dùng auxiliary array
// Radix Sort: sắp xếp theo từng digit. Time O(w·n), thích hợp với khóa có độ dài cố định, không dùng so sánh
// Bucket Sort: chia dữ liệu thành các bucket rồi sort từng bucket (có thể dùng các sorts khác). Hiệu quả nếu dữ liệu đồng đều phân bố

// ! Hybrid sorting
// Hybrid sorts kết hợp nhiều thuật toán để tận dụng điểm mạnh và giảm nhược điểm:

// 	? Introsort: bắt đầu Quick Sort, nhưng nếu đệ quy quá sâu, chuyển sang Heap Sort để tránh worst-case O(n²).
// 	? TimSort: như trên, phổ biến trong Python, Java,...
// 	? Spread sort, Block sort: kết hợp Radix + Quick hoặc Merge + Insertion

/*
	! What is "stable" in sorting?
	- Stable sort: nếu hai phần tử A và B có key bằng nhau, sau khi sắp xếp phần tử nào đứng trước ban đầu vẫn đứng trước ở kết quả.
	- Unstable sort: không đảm bảo giữ thứ tự tương đối của các phần tử có key bằng nhau — thứ tự đó có thể đổi.

	? Merge Sort, Timsort, Counting Sort (stable implementations) → stable

	? QuickSort, HeapSort → thường unstable.
*/
