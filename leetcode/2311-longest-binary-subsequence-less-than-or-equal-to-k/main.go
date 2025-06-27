package main

import (
	"fmt"
	"strconv"
)

// leetcode: https://leetcode.com/problems/longest-binary-subsequence-less-than-or-equal-to-k/

func longestSubsequence(s string, k int) int {
	for i := len(s) - 1; i >= 0; i-- {
		for j := i; j >= 0; j-- {

		}
	}
}

func main() {

	// Example usage
	s := "1001010"
	k := 5
	result := longestSubsequence(s, k)
	fmt.Println("Longest binary subsequence length:", result)

	fmt.Println("Hello")
}
