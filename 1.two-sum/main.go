package main

import "fmt"

func hashMapSolution(nums []int, target int) []int {
	var mapNums = make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if j, ok := mapNums[target-nums[i]]; ok {
			return []int{j, i}
		}
		mapNums[nums[i]] = i
	}
	return nil
}

func bruteForce(nums []int, target int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return nil
}

func main() {
	var nums = []int{1, 2, 4, 6, 7}
	var target = 6
	fmt.Println(hashMapSolution(nums, target))
	fmt.Println(bruteForce(nums, target))

}
