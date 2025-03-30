package main

import "fmt"

func partitionLabels(s string) []int {
	lastOccurrence := make(map[rune]int)
	result := make([]int, 0)

	for i, char := range s {
		lastOccurrence[char] = i
	}

	start, end := 0, 0

	for i, char := range s {
		end = max(end, lastOccurrence[char])
		if i == end {
			result = append(result, end-start+1)
			start = i + 1
		}
	}

	return result
}

func main() {
	str := "ababcbacadefegdehijhklij"
	fmt.Println(partitionLabels(str))
}
