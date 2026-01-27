package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	mergeSortedArr := make([]int, m+n)

	i, j := 0, 0
	k := 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] < nums2[j] {
			mergeSortedArr[k] = nums1[i]
			i++
		} else {
			mergeSortedArr[k] = nums2[j]
			j++
		}
		k++
	}

	for ; i < m; i++ {
		mergeSortedArr[k] = nums1[i]
		k++
	}

	for ; j < n; j++ {
		mergeSortedArr[k] = nums2[j]
		k++
	}

	//cal median
	sum := (m + n)
	if sum%2 == 0 {
		mid := sum/2 - 1
		return (float64(mergeSortedArr[mid]) + float64(mergeSortedArr[mid+1])) / 2.0
	}

	mid := (sum+1)/2 - 1
	return float64(mergeSortedArr[mid])
}

func main() {
	nums1 := []int{1, 2}
	nums2 := []int{3, 4}
	res := findMedianSortedArrays(nums1, nums2)
	fmt.Println(res)
}
