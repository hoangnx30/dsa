package main

import "fmt"

func maximumTripletValueN3(nums []int) int64 {
	n := len(nums)
	maxValue := 0
	for i := 0; i < n-2; i++ {
		for j := i + 1; j < n-1; j++ {
			for k := j + 1; k < n; k++ {
				maxValue = max(maxValue, (nums[i]-nums[j])*nums[k])
			}
		}
	}

	return int64(max(maxValue, 0))
}

func maximumTripletValueN(nums []int) int64 {
	length := len(nums)
	maxValue := 0

	prefixMax := make([]int, length)
	prefixMax[0] = nums[0]
	for i := 1; i < length; i++ {
		prefixMax[i] = max(prefixMax[i-1], nums[i])
	}

	suffixMax := make([]int, length)
	suffixMax[length-1] = nums[length-1]
	for i := length - 2; i >= 0; i-- {
		suffixMax[i] = max(suffixMax[i+1], nums[i])
	}

	fmt.Println(prefixMax, suffixMax)

	for i := 1; i < length-1; i++ {
		maxValue = max(maxValue, (prefixMax[i-1]-nums[i])*suffixMax[i+1])
	}

	return int64(max(maxValue, 0))
}

func main() {
	nums := []int{2, 3, 1}

	fmt.Println(maximumTripletValueN(nums))
}
