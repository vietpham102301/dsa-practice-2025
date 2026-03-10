package main

import "fmt"

func moveZeroes(nums []int) {
	if len(nums) <= 1 {
		return
	}

	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}
}

func main() {
	test := []int{0, 1, 0, 3, 12}

	moveZeroes(test)
	fmt.Println(test)
}
