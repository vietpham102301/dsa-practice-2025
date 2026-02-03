package main

import "fmt"

func reverseArr(nums []int) {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	reverseArr(nums)
	fmt.Println(nums)
}
