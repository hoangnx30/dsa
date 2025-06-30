package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	// i := m
	// j := 0
	// for i < m + n {
	//   nums1[i] = nums2[j]
	//   i++
	//   j++
	// }

	// sort.Ints(nums1)

	// 4 5 6 0 0 0
	// 1 2 3

	// 1 2 3 0 0 0
	// 2 5 6

	right := m + n - 1
	idx1 := m - 1
	idx2 := n - 1

	for idx2 >= 0 {
		fmt.Printf("idx1: %d, idx2: %d, right: %d\n", idx1, idx2, right)
		if idx1 >= 0 && nums1[idx1] > nums2[idx2] {
			nums1[right] = nums1[idx1]
			idx1--

		} else {
			nums1[right] = nums2[idx2]
			idx2--
		}

		right--
	}
}

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	m := 3
	n := 3
	merge(nums1, m, nums2, n)
	fmt.Println(nums1) // Output: [1, 2, 2, 3, 5, 6]
	// Output: [1, 2, 2, 3, 5, 6]

	nums1 = []int{0, 0, 0, 0, 0}
	nums2 = []int{1, 2, 3, 4, 5}
	m = 0
	n = 5
	merge(nums1, m, nums2, n)
	fmt.Println(nums1) // Output: [1, 2, 3, 4, 5, 6]

}
