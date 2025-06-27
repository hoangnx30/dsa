package main

import "fmt"

func findPermutation(s string) []string {
	var result []string
	n := len(s)
	used := make([]bool, n)
	currentPermutation := make([]rune, 0, n)

	var dfs func()
	dfs = func() {
		if len(currentPermutation) == n {
			result = append(result, string(currentPermutation))
			return
		}

		for i := 0; i < n; i++ {
			if !used[i] {
				used[i] = true
				currentPermutation = append(currentPermutation, rune(s[i]))

				dfs()

				currentPermutation = currentPermutation[:len(currentPermutation)-1]
				used[i] = false
			}
		}

	}
	dfs()

	return result
}

func main() {
	inputString := "abca"
	permutations := findPermutation(inputString)
	fmt.Printf("Input string: '%s'\n", inputString)
	fmt.Printf("All permutations: %v\n", permutations)

	inputString2 := "ab"
	permutations2 := findPermutation(inputString2)
	fmt.Printf("Input string: '%s'\n", inputString2)
	fmt.Printf("All permutations: %v\n", permutations2)
}
