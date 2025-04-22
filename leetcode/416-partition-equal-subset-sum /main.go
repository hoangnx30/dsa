package main

import "fmt"

func canPartition(nums []int) bool {
	totalSum := 0
	for _, num := range nums {
		totalSum += num
	}

	// 1. If the total sum is odd, it's impossible to partition into two equal subsets.
	if totalSum%2 != 0 {
		return false
	}

	targetSum := totalSum / 2

	// 2. Initialize DP slice.
	// dp[j] will be true if sum j can be achieved using a subset of numbers seen so far.
	// The slice is automatically initialized with false values.
	// Size is targetSum + 1 to include sum 0 up to targetSum.
	dp := make([]bool, targetSum+1)
	dp[0] = true // Base case: sum 0 is always achievable with an empty subset

	// 3. Iterate through numbers and update DP table
	for _, num := range nums {
		// Iterate downwards from targetSum down to the current number (num).
		// This prevents using the same element multiple times within one subset calculation
		// for a specific target sum j in this iteration.
		for j := targetSum; j >= num; j-- {
			// If sum j was already possible OR sum (j - num) was possible before
			// considering the current 'num', then sum j is now possible (by potentially adding 'num').
			dp[j] = dp[j] || dp[j-num]
		}
		// Optional optimization check (can sometimes return early):
		// if dp[targetSum] {
		//     return true
		// }

		fmt.Println(dp)
	}

	// 4. The final answer is whether the targetSum is achievable
	return dp[targetSum]
}

func main() {
	// --- Example Usage ---
	//fmt.Printf("nums = [1, 5, 11, 5], Result: %t\n", canPartition([]int{1, 5, 11, 5}))     // Output: true
	//fmt.Printf("nums = [1, 2, 3, 5], Result: %t\n", canPartition([]int{1, 2, 3, 5}))       // Output: false
	//fmt.Printf("nums = [2, 2, 3, 5], Result: %t\n", canPartition([]int{2, 2, 3, 5}))       // Output: false
	//fmt.Printf("nums = [1, 2, 5], Result: %t\n", canPartition([]int{1, 2, 5}))             // Output: false
	//fmt.Printf("nums = [3, 3, 3, 4, 5], Result: %t\n", canPartition([]int{3, 3, 3, 4, 5})) // Output: true
	//fmt.Printf("nums = [100], Result: %t\n", canPartition([]int{100}))                     // Output: false
	//fmt.Printf("nums = [1,1], Result: %t\n", canPartition([]int{1, 1}))                    // Output: true
	fmt.Printf("nums = [3, 3, 4, 2], Result: %t\n", canPartition([]int{6, 3, 3, 2, 0})) // Output: true

}
