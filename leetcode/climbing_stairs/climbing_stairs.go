package main

import "fmt"

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}

	if n == 2 {
		return 2
	}

	step1 := 1 //clibm to n-2
	step2 := 2 //climb to n-1
	res := 0
	for i := 3; i <= n; i++ {
		//number of step to climb to n-1
		res = step1 + step2
		//number of step to climb to n-2
		step1 = step2
		step2 = res
	}

	return res
}

func main() {
	n := 5
	res := climbStairs(n)

	fmt.Printf("res: %d", res)
}
