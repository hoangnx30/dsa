package main

import "fmt"

// use map to store the lastIndex of characters
func lengthOfLongestSubstring(s string) int {
	maxLength := 0
	start := 0
	charMap := make(map[byte]int)
	for end := range len(s) {
		// update the start value when it's current window only.
		if pos, ok := charMap[s[end]]; ok && pos >= start {
			start = pos + 1
		}

		charMap[s[end]] = end
		maxLength = max(maxLength, end-start+1)
	}

	return maxLength
}

// using array to keep last index of each character
func solution2(s string) int {
	lastIndex := make([]int, 256)
	for i := range lastIndex {
		lastIndex[i] = -1
	}

	maxLength := 0
	start := 0
	for end := range len(s) {
		if lastIndex[s[end]] >= start {
			start = lastIndex[s[end]] + 1
		}

		lastIndex[s[end]] = end
		maxLength = max(maxLength, end-start+1)
	}

	return maxLength
}

func bruteForce(s string) int {
	n := len(s)
	maxLength := 0
	// Try all possible substrings
	for start := 0; start < n; start++ {
		for end := start; end < n; end++ {
			if hasUniqueChar(s, start, end) {
				maxLength = max(maxLength, end-start+1)
			}
		}
	}

	return maxLength
}

// Helper function to check if a substring has all unique characters
func hasUniqueChar(s string, start int, end int) bool {
	charMap := make(map[byte]bool)

	for i := start; i <= end; i++ {
		if charMap[s[i]] {
			return false
		}
		charMap[s[i]] = true
	}

	return true
}

func main() {
	//s := "abcabcbb"
	s := "abbaack"

	fmt.Println("Solution 1: ", lengthOfLongestSubstring(s))
	fmt.Println("Solution 2: ", solution2(s))
	fmt.Println("Solution 3: ", bruteForce(s))
}
