package main

import "fmt"

func mergeArray(nums1 []int, nums2 []int) []int {
	nums1Index, nums2Index := 0, 0
	result := make([]int, 0, len(nums1)+len(nums2))
	for nums1Index < len(nums1) && nums2Index < len(nums2) {
		if nums1[nums1Index] < nums2[nums2Index] {
			result = append(result, nums1[nums1Index])
			nums1Index++
		} else {
			result = append(result, nums2[nums2Index])
			nums2Index++
		}
	}
	result = append(result, nums1[nums1Index:]...)
	result = append(result, nums2[nums2Index:]...)
	return result
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	finalArray := mergeArray(nums1, nums2)
	fmt.Println(finalArray)

	if len(finalArray)%2 != 0 {
		return float64(finalArray[len(finalArray)/2])
	}
	return float64(finalArray[len(finalArray)/2]+finalArray[len(finalArray)/2-1]) / 2

}

func main() {
	arr1 := []int{1, 3, 5, 7, 9}
	arr2 := []int{2, 4, 6, 8, 10}
	fmt.Println(findMedianSortedArrays(arr1, arr2))
}
