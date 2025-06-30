package main

import (
	"fmt"
	"math"
	"sort"
)

func numSubseq(nums []int, target int) int {
	mod := int(math.Pow(10, 9) + 7)
	result := 0
	power := make([]int, len(nums))
	power[0] = 1
	for i := 1; i < len(nums); i++ {
		power[i] = (power[i-1] * 2) % mod
	}
	// Sort the array to use two pointers
	// This is important to ensure that we can find pairs efficiently
	// and to ensure that we can use the two-pointer technique
	// to find the valid subsequences
	sort.Ints(nums)

	left, right := 0, len(nums)-1
	for left <= right {
		if nums[left]+nums[right] > target {
			right-- // If the sum is greater than target, move the right pointer left
		} else {
			// If the sum is less than or equal to target,
			// all subsequences formed by nums[left] and any of the elements
			// between left and right (inclusive) are valid
			result = (result + power[right-left]) % mod
			left++ // Move the left pointer right to consider the next element
		}
	}
	return result
}

func main() {
	nums := []int{3, 5, 6, 7}
	target := 9

	result := numSubseq(nums, target)
	// Example usage
	fmt.Printf("Number of subsequences that satisfy the given sum condition with %v: %d\n", nums, result)

	nums = []int{3, 3, 6, 8}
	target = 10
	result = numSubseq(nums, target)
	fmt.Printf("Number of subsequences that satisfy the given sum condition with %v: %d\n", nums, result)

}
