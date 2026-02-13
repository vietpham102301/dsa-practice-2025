package main

import (
	"fmt"

	"github.com/vietpham102301/dsa-practice-2025/sorting_algorithms/algorithms"
)

func main() {
	unsortedArray := []int{0, 1, 2, -2, -1}

	sortedArray := algorithms.HeapSort(unsortedArray)
	fmt.Println(sortedArray)
}
