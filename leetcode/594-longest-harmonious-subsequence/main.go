package main

import "fmt"

func findLHS(nums []int) int {
	freq := make(map[int]int)
	max := 0
	for _, val := range nums {
		freq[val]++
	}

	for _, val := range nums {
		if _, existed := freq[val+1]; existed {
			newMax := freq[val+1] + freq[val]
			if newMax > max {
				max = newMax
			}
		}
	}

	return max
}

func main() {
	nums := []int{1, 3, 2, 2, 5, 4, 3, 7, 8, 6}
	result := findLHS(nums)
	fmt.Println("Longest Harmonious Subsequence Length:", result) // Expected output: 5
}
