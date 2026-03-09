package main

func permute(nums []int) [][]int {
	res := [][]int{}
	used := make(map[int]bool)
	findPermute(nums, []int{}, &res, &used)
	return res
}

func findPermute(nums []int, path []int, res *[][]int, used *map[int]bool) {
	n := len(nums)

	if len(path) == n {
		temp := make([]int, n)
		copy(temp, path)
		(*res) = append((*res), temp)
		return
	}

	for i := 0; i < n; i++ {
		if (*used)[i] {
			continue
		}

		(*used)[i] = true
		path = append(path, nums[i])
		findPermute(nums, path, res, used)

		//backtrack: try other option
		path = path[:len(path)-1]
		(*used)[i] = false
	}
}
