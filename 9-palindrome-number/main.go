package main

import "fmt"

// isPalindrome checks if an integer is a palindrome.
func isPalindrome(x int) bool {
	xCopy := x
	reversedNum := 0
	for xCopy > 0 {
		reversedNum = (reversedNum * 10) + xCopy%10
		xCopy = xCopy / 10
	}

	return x == reversedNum

}

func main() {
	num := 123
	fmt.Println(isPalindrome(num))
}

