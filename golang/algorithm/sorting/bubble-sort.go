package main

/*
	5, 2, 4, 6, 1, 3
	2, 4, 5, 1, 3, 6
	2, 4, 5, 6, 1, 3
*/

func BubbleSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j+1] < arr[j] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}
