package main

import "fmt"

func binSearch(arr []int, target int) bool {
	left_idx := 0
	right_idx := len(arr)
	for left_idx < right_idx {
		mid_idx := (left_idx + right_idx) / 2
		if arr[mid_idx] == target {
			return true
		} else if arr[mid_idx] < target {
			left_idx = mid_idx + 1
		} else {
			right_idx = mid_idx
		}
	}
	return false
}

func main() {
	arr := []int{5, 6, 7, 8, 9, 10}
	fmt.Println(binSearch(arr, 10))
}
