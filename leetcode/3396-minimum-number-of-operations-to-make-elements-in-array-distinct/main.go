package main

import (
	"fmt"
)

func minOperationReverse(nums []int) int {
	seen := make(map[int]bool)
	startDistinctIndex := 0
	for i := len(nums) - 1; i >= 0; i-- {
		if _, ok := seen[nums[i]]; ok {
			startDistinctIndex = i + 1
			break
		}
		seen[nums[i]] = true
	}

	operations := (startDistinctIndex + 2) / 3
	return operations
}

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

func runTests(minOpsFunc func([]int) int) {
	test1 := []int{1, 2, 3, 3, 1, 2, 3, 4, 5}
	fmt.Printf("Input: %v, Min Operations: %d\n", test1, minOpsFunc(test1)) // Expected: 2

	test2 := []int{1, 1, 1, 2, 2, 2}
	fmt.Printf("Input: %v, Min Operations: %d\n", test2, minOpsFunc(test2)) // Expected: 2

	test3 := []int{1, 2, 3, 4, 5}
	fmt.Printf("Input: %v, Min Operations: %d\n", test3, minOpsFunc(test3)) // Expected: 0

	test4 := []int{1, 1, 2, 2}
	fmt.Printf("Input: %v, Min Operations: %d\n", test4, minOpsFunc(test4)) // Expected: 1

	test5 := []int{}
	fmt.Printf("Input: %v, Min Operations: %d\n", test5, minOpsFunc(test5)) // Expected: 0

	test6 := []int{3, 3, 3, 3, 3, 3}
	fmt.Printf("Input: %v, Min Operations: %d\n", test6, minOpsFunc(test6)) // Expected: 2

	test7 := []int{1, 2, 1, 2, 1, 2}
	fmt.Printf("Input: %v, Min Operations: %d\n", test7, minOpsFunc(test7)) // Expected: 2

}

func main() {
	fmt.Println("--- Testing Direct Simulation ---")
	runTests(minimumOperations) // Using the function from the previous answer

	fmt.Println("--- Testing Recursion Simulation ---")
	runTests(minOperationReverse) // Using the function from the previous answer

	fmt.Println((2 + 2) / 3)

}
