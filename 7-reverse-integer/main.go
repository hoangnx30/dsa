package main

import (
	"fmt"
	"math"
	"strconv"
)

func reverse(num int) int {
	isNegative := num < 0

	numString := strconv.Itoa(int(math.Abs(float64(num))))

	runes := []rune(numString)

	for i := 0; i < len(runes)/2; i++ {
		runes[i], runes[len(numString)-1-i] = runes[len(numString)-1-i], runes[i]
	}
	result, _ := strconv.Atoi(string(runes))

	if isNegative {
		result = -result
	}

	if math.MinInt32 > result || result > math.MaxInt32 {
		return 0
	}

	return result
}

func main() {
	//num := 312
	num := 1534236469
	fmt.Println(reverse(num))

}
