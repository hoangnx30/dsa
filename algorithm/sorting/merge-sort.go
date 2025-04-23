package main

func MergeSort(arr []int) []int {
	if len(arr) == 1 {
		return arr
	}

	mid := len(arr) / 2

	left := MergeSort(arr[0:mid])
	right := MergeSort(arr[mid:])

	return mergeTwoSortedArray(left, right)
}

func mergeTwoSortedArray(arr1 []int, arr2 []int) []int {
	res := make([]int, 0, len(arr1)+len(arr2))
	i, j := 0, 0
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			res = append(res, arr1[i])
			i++
		} else {
			res = append(res, arr2[j])
			j++
		}
	}
	res = append(res, arr1[i:]...)
	res = append(res, arr2[j:]...)
	return res
}
