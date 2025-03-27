package main

import "fmt"

func findMaxAverage(nums []int, k int) float64 {
	maxSum := -999
	start := 0
	sum := 0
	for end := 0; end < len(nums); end++ {
		if end-start+1 < k {
			sum += nums[end]
		} else {
			sum += nums[end]
			maxSum = max(maxSum, sum)

			sum -= nums[start]
			start++
		}
	}

	return float64(maxSum) / float64(k)
}

func main() {
	arr := []int{1, 12, -5, -6, 50, 3}
	k := 4
	result := findMaxAverage(arr, k)

	fmt.Println(result)
}
