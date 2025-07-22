package main

import "fmt"

func isValid(s string) bool {
	stack := []rune{}
	pairs := map[rune]rune{']': '[', '}': '{', ')': '('}

	for _, val := range s {
		if val == '(' || val == '[' || val == '{' {
			stack = append(stack, val)
		} else {
			n := len(stack) - 1

			if n < 0 || stack[n] != pairs[val] {
				return false
			}

			if stack[n] == pairs[val] {
				stack = stack[:n]
			}
		}
	}

	return len(stack) == 0
}

func main() {
	s := "()[]{}"
	result := isValid(s)
	fmt.Println("Is valid parentheses:", result) // Expected output: true

	s = "([])"
	result = isValid(s)
	fmt.Println("Is valid parentheses:", result) // Expected output: true

	s = "]"
	result = isValid(s)
	fmt.Println("Is valid parentheses:", result) // Expected output: true
}
