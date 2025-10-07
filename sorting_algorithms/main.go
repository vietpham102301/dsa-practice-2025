package main

import (
	"fmt"
	"github.com/vietpham102301/dsa-practice-2025/sorting_algorithms/merge_sort"
)

func main() {
	unsortedArray := []int{3, 1, 5, 2, 3, 7, 9}

	sortedArray := merge_sort.MergeSort(unsortedArray) //O(nlogn)

	fmt.Println(sortedArray)
}
