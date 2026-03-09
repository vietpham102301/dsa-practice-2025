package main

import "fmt"

func subsets(nums []int) [][]int {
	res := make([][]int, 0)
	path := []int{}
	findSubset(nums, 0, path, &res)

	return res
}

func findSubset(nums []int, start int, path []int, res *[][]int) {
	temp := make([]int, len(path))
	copy(temp, path)

	(*res) = append((*res), temp)

	for i := start; i < len(nums); i++ {
		path = append(path, nums[i])
		findSubset(nums, i+1, path, res)

		path = path[:len(path)-1]
	}
}

func subsetsIterative(nums []int) [][]int {
	res := [][]int{{}}
	for _, num := range nums {
		n := len(res)
		for i := 0; i < n; i++ {
			temp := make([]int, len(res[i]))
			copy(temp, res[i])
			temp = append(temp, num)

			res = append(res, temp)
		}
	}

	return res
}

func main() {
	nums := []int{1, 2, 3}
	res := subsets(nums)
	fmt.Println(res)
}
