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

func main() {
	var nums = []int{1, 2, 4, 6, 7}
	var target = 6
	fmt.Println(hashMapSolution(nums, target))

}
