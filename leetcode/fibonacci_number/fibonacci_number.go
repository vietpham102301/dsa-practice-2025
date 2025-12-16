package main

import "fmt"

func fib(n int) int {
	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	return fib(n-1) + fib(n-2)
}

func fibDPTopDown(n int, memos []int) int {
	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	if memos[n] != 0 {
		return memos[n]
	}

	res := fibDPTopDown(n-1, memos) + fibDPTopDown(n-2, memos)
	memos[n] = res
	return res
}

func fibDPBottomUp(n int) int {
	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	first, second := 0, 1
	res := 0
	for i := 2; i <= n; i++ {
		res = first + second
		first = second
		second = res
	}

	return res
}

func main() {
	n := 5
	memos := make([]int, 30)
	fmt.Printf("fib n: %d", fibDPTopDown(n, memos))
}
