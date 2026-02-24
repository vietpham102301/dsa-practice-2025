package main

import "fmt"

func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	stack := make([]int, 0, n)
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		for len(stack) > 0 && temperatures[stack[len(stack)-1]] < temperatures[i] {
			l := len(stack) - 1
			idx := stack[l]
			stack = stack[:l]
			ans[idx] = i - idx
		}

		stack = append(stack, i)
	}

	return ans
}

func main() {
	temperatures := []int{73, 74, 75, 71, 69, 72, 76, 73}

	res := dailyTemperatures(temperatures)
	fmt.Println(res)
}
