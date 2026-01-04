package main

import "fmt"

func shipWithinDays(weights []int, days int) int {
	left, right := findMinCap(weights), findMaxCap(weights)
	ans := -1
	for left <= right {
		mid := (left + right) / 2
		if canShip(weights, mid, days) {
			ans = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return ans
}

func findMaxCap(weights []int) int {
	s := 0
	for _, w := range weights {
		s += w
	}

	return s
}
func findMinCap(weights []int) int {
	s := 0
	for _, w := range weights {
		if w > s {
			s = w
		}
	}

	return s
}

func canShip(weights []int, cap int, d int) bool {
	l := 0
	days := 1
	for _, w := range weights {
		l += w
		if l > cap {
			days += 1
			l = w
		}
	}

	return days <= d
}

func main() {
	weights := []int{1, 2, 3, 1, 1}
	days := 4

	res := shipWithinDays(weights, days)
	fmt.Println(res)
}
