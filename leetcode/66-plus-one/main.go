package main

import (
	"fmt"
)

func plusOne(digits []int) []int {
	remind := 0

	for i := len(digits) - 1; i >= 0; i-- {
		newDigit := digits[i]

		if i == len(digits)-1 {
			newDigit++
		}

		newDigit += remind

		if newDigit == 10 {
			digits[i] = 0
			remind = 1
		} else {
			digits[i] = newDigit
			remind = 0
		}
	}

	if remind > 0 {
		return append([]int{1}, digits...)
	}

	return digits

}

func main() {
	digits := []int{9}
	fmt.Println(plusOne(digits))

	digits = []int{1, 2, 3}
	fmt.Println(plusOne(digits))

}
