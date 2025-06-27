package main

import "fmt"

func removeElement(nums []int, val int) int {
	res := 0
	for i := range nums {
		if nums[i] != val {
			nums[res] = val
			res++
		}
	}
	return res
}

func main() {
	nums := []int{3, 3, 2, 2}
	val := 3

	fmt.Println(removeElement(nums, val))
}

