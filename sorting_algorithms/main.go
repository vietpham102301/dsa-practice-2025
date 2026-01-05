package main

import (
	"fmt"

	"github.com/vietpham102301/dsa-practice-2025/sorting_algorithms/algorithms"
)

func main() {
	unsortedArray := []int{3, 1, 5, 2, 3, 7, 9}

	sortedArray := algorithms.InsertionSort(unsortedArray)
	fmt.Println(sortedArray)
}
