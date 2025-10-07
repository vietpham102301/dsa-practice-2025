package main

import "fmt"

func mergeSortedArr(nums1 []int, m int, nums2 []int, n int) {
	l := cap(nums1) - 1
	i, j := m-1, n-1

	for l >= 0 {
		if j < 0 {
			nums1[l] = nums1[i]
			i--
		} else if i < 0 {
			nums1[l] = nums2[j]
			j--
		} else if nums1[i] < nums2[j] {
			nums1[l] = nums2[j]
			j--
		} else {
			nums1[l] = nums1[i]
			i--
		}
		l--
	}
}

func main() {
	nums1 := []int{1, 3, 5, 0, 0, 0}
	m := 3
	nums2 := []int{2, 4, 6}
	n := 3

	mergeSortedArr(nums1, m, nums2, n)
	fmt.Println(nums1) // → [1 2 3 4 5 6]
}
