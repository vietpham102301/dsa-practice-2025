package main

import (
	"fmt"
)

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}

	neg := n < 0
	if neg {
		n = -n
	}

	res := 1.0
	for n > 0 {
		if n&1 == 1 { //chọn các bit là 1 để tích luỹ vào res
			res *= x
		}
		x *= x
		n >>= 1
	}

	if neg {
		return 1 / res
	}

	return res
}

func myPowRecursion(x float64, n int) float64 {
	if n == 0 {
		return 1
	}

	if n < 0 {
		return 1 / myPowRecursion(x, -n)
	}

	if n%2 == 0 {
		half := myPowRecursion(x, n/2)
		return half * half
	}

	return x * myPowRecursion(x, n-1)
}

func main() {
	x := 2.0
	n := 10

	res := myPowRecursion(x, n)
	fmt.Printf("res: %f", res)
}
