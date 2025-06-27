package main

import "fmt"

/*
- 3 2 5 1 4
- 3 2 5 1 4 (left = 1, right =  4, pivot = arr[low], nextPivotPosition = 1)
- 3 2 5 1 4 (left = 2, right =  4, pivot = arr[low], nextPivotPosition = 1)
- 3 2 1 5 4 (left = 3, right =  4, pivot = arr[low], nextPivotPosition = 2)
- 3 2 5 5 4 (left = 4, right =  4, pivot = arr[low], nextPivotPosition = 2)
- 1 2 3 5 4

- 1 2 3 5 4
- 1 2       (left = 0, right =  1, pivot = arr[low], nextPivotPosition = 0)
- 1 2       (left = 0, right =  1, pivot = arr[low], nextPivotPosition = 0)

- 1 2       (left = 0, right = -1, pivot = arr[low], nextPivotPosition = 0) // end
- 1 2       (left = 1, right = 1, pivot = arr[low], nextPivotPosition = 0) // end

-       5 4 (left = 3, right =  4,  pivot = arr[low], nextPivotPosition = 3)
-       5 4 (left = 4, right =  4,  pivot = arr[low], nextPivotPosition = 4)
-       4 5
-           (left = 3, right =  3,  pivot = arr[low]) // end
-           (left = 5, right =  4,  pivot = arr[low]) // end
*/
func QuickSort(arr []int) {
	_quickSort(arr, 0, len(arr)-1)
}

func _quickSort(arr []int, left int, right int) {
	fmt.Println(left, right)
	if left < right {
		pivot := partition(arr, left, right)
		_quickSort(arr, left, pivot-1)
		_quickSort(arr, pivot+1, right)

	}
}

func partition(arr []int, left int, right int) int {
	pivot := arr[left]
	nextPivotPosition := left
	for i := left + 1; i <= right; i++ {
		if arr[i] <= pivot {
			nextPivotPosition++
			arr[nextPivotPosition], arr[i] = arr[i], arr[nextPivotPosition]
		}
	}
	arr[left], arr[nextPivotPosition] = arr[nextPivotPosition], arr[left]
	return nextPivotPosition

}
