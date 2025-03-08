package main

import "fmt"

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

func main() {
	//s := "abcabcbb"
	s := "abbaack"

	fmt.Println(lengthOfLongestSubstring(s))
}
