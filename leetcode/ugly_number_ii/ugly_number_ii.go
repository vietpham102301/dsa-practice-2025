package main

import "fmt"

func nthUglyNumber(n int) int {
	dp := make([]int, n)

	p2 := 0
	p3 := 0
	p5 := 0

	dp[0] = 1

	for i := 1; i < n; i++ {
		next2 := dp[p2] * 2
		next3 := dp[p3] * 3
		next5 := dp[p5] * 5

		dp[i] = min(next2, min(next3, next5))
		if next2 == dp[i] {
			p2++
		}
		if next3 == dp[i] {
			p3++
		}
		if next5 == dp[i] {
			p5++
		}
	}

	return dp[n-1]
}

func main() {
	n := 10
	res := nthUglyNumber(n)
	fmt.Println(res)
}
