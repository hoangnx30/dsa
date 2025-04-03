package main

import "fmt"

func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2

	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])

	return merge(left, right)
}

func merge(left []int, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	leftIndex, rightIndex := 0, 0
	// [1,5], [3,4]

	for leftIndex < len(left) && rightIndex < len(right) {
		if left[leftIndex] < right[rightIndex] {
			result = append(result, left[leftIndex])
			leftIndex++
		} else {
			result = append(result, right[rightIndex])
			rightIndex++
		}
	}

	// Add remaining elements from left array, if any
	result = append(result, left[leftIndex:]...)
	// Add remaining elements from right array, if any
	result = append(result, right[rightIndex:]...)

	return result
}

func main() {
	arr := []int{5, 62, 1, 3, 676, 342, 23, 25}

	fmt.Println(mergeSort(arr))
}
