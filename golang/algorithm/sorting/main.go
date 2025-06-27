package main

import "fmt"

func main() {
	arr := []int{3, 2, 5, 1, 4}
	// SelectionSort(arr)
	// fmt.Println("Sorted array with selection sort", arr)
	// BubbleSort(arr)
	// fmt.Println("Sorted array with bubble sort", arr)
	// fmt.Println("Sorted array with merge sort", MergeSort(arr))
	QuickSort(arr)
	fmt.Println("Sorted array with quick sort", arr)

}
