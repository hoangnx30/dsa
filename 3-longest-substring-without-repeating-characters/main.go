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

func main() {
	//s := "abcabcbb"
	s := "abbaack"

	fmt.Println("Solution 1: ", lengthOfLongestSubstring(s))
	fmt.Println("Solution 2: ", solution2(s))
}
