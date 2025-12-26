package main

import "fmt"

func searchRange(nums []int, target int) []int {
	return []int{firstPost(nums, target), lastPost(nums, target)}
}

func firstPost(nums []int, target int) int {
	l, r := 0, len(nums)-1
	ans := -1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] >= target {
			r = mid - 1
		} else {
			l = mid + 1
		}
		if nums[mid] == target {
			ans = mid
		}
	}

	return ans
}

func lastPost(nums []int, target int) int {
	l, r := 0, len(nums)-1
	ans := -1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] <= target {
			l = mid + 1
		} else {
			r = mid - 1
		}
		if nums[mid] == target {
			ans = mid
		}
	}

	return ans
}

func main() {
	nums := []int{5, 7, 7, 8, 8, 10}
	target := 8
	res := searchRange(nums, target)
	fmt.Println(res)
}
