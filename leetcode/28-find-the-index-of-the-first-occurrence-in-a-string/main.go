package main

import "fmt"

func strStr(haystack string, needle string) int {
	for i := 0; i <= len(haystack)-len(needle); i++ {
		temp := haystack[i : i+len(needle)]
		fmt.Println(temp)
		if temp == needle {
			return i
		}
	}
	return -1
}

func main() {
	haystack := "a"
	needle := "a"
	result := strStr(haystack, needle)
	fmt.Printf("Result: %d\n", result)
}
