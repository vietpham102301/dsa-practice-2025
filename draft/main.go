package main

import "fmt"

func totalDays(weights []int, cap int) int {
	l := 0
	days := 1
	for _, w := range weights {
		l += w
		if l > cap {
			days += 1
			l = w
		}
	}

	return days
}

func main() {
	weights := []int{1, 2, 3, 1, 1}
	cap := 3

	days := totalDays(weights, cap)
	fmt.Println(days)
}
