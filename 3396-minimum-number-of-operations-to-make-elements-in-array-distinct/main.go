package main

import "fmt"

func hasDuplicatedElements(nums []int) bool {
	seen := make(map[int]bool)

	for _, num := range nums {
		if _, existed := seen[num]; existed {
			return true
		}
		seen[num] = true
	}

	return false
}

func minimumOperations(nums []int) int {
	numsLen := len(nums)
	operation := 0
	for {
		startIndex := 3 * operation

		if startIndex >= numsLen {
			return operation
		}

		subArray := nums[startIndex:]

		if !hasDuplicatedElements(subArray) {
			return operation
		}

		operation++

	}
}

func main() {
	// --- Example Usage ---
	//fmt.Printf("nums = [1,2,3,4,2,3,3,5,7], Result: %t\n", minimumOperations([]int{1, 2, 3, 4, 2, 3, 3, 5, 7})) // Output: 2
	//fmt.Printf("nums = [4,5,6,4,4], Result: %t\n", minimumOperations([]int{4, 5, 6, 4, 4}))                     // Output: 2
	fmt.Printf("nums = [6, 7, 8, 9], Result: %t\n", minimumOperations([]int{6, 7, 8, 9})) // Output: 0

}
