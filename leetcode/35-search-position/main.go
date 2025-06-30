package main

func searchInsert(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := (left + right) / 2

		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right--
		} else {
			left++
		}
	}

	return left
}

func main() {
	nums := []int{1, 3, 5, 6}
	target := 5
	result := searchInsert(nums, target)
	println("Search Insert Position:", result) // Expected output: 2
	target = 2
	result = searchInsert(nums, target)
	println("Search Insert Position:", result) // Expected output: 1
}
