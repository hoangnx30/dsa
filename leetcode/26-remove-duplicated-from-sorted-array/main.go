package main

import "fmt"

func removeDuplicates(nums []int) int {
	count := 1
	idx := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[idx] = nums[i]
			count++
			idx++
		}
	}
	return count
}

func main() {
	nums := []int{1, 1, 2}
	result := removeDuplicates(nums)

	fmt.Println(result)
}

