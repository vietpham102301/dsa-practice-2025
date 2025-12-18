package main

import "fmt"

func isPerfectSquare(num int) bool {
	if num == 1 {
		return true
	}

	left, right := 1, num/2

	for left <= right {
		mid := (left + right) / 2

		s := mid * mid
		if s == num {
			return true
		} else if s < num {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return false
}

func main() {
	num := 9

	res := isPerfectSquare(num)
	fmt.Println(res)
}
