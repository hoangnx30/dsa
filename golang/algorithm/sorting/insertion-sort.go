package main

/*
*
5, 2, 4, 6, 1, 3
2, 5, 4, 6, 1, 3
2, 4, 5, 6, 1, 3
*/
func InsertionSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		current := arr[i]
		j := i - 1
		for j >= 0 && current < arr[j] {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = current

	}
}
