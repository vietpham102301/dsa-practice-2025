package main

import "fmt"

func minEatingSpeed(piles []int, h int) int {
	left, right := 1, findMaxSpeed(piles)
	ans := 1
	for left <= right {
		mid := (left + right) / 2
		t := totalTime(piles, mid)
		if t <= h {
			ans = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return ans
}

func totalTime(piles []int, k int) int {
	s := 0
	for _, p := range piles {
		s += ceil(p, k)
	}

	return s
}

func findMaxSpeed(piles []int) int {
	m := 1
	for _, p := range piles {
		if p > m {
			m = p
		}
	}

	return m
}

func ceil(p, k int) int {
	return (p + k - 1) / k
}

func main() {
	piles := []int{3, 6, 7, 11}
	h := 8
	res := minEatingSpeed(piles, h)

	fmt.Println(res)
}
