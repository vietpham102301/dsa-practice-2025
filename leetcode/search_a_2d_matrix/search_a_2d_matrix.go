package main

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}

	m, n := len(matrix), len(matrix[0])

	var flatternArr []int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			flatternArr = append(flatternArr, matrix[i][j])
		}
	}

	left, right := 0, len(flatternArr)-1
	for left <= right {
		mid := (left + right) / 2
		if flatternArr[mid] == target {
			return true
		}

		if target > flatternArr[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return false
}

func main() {
	matrix := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}
	target := 3
	res := searchMatrix(matrix, target)
	fmt.Println(res)
}
