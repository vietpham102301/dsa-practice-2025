package main

import "fmt"

func GCDEuclidean(a, b int) int {
	if b == 0 {
		if a < 0 {
			return -a // GCD(a, 0) = |a|
		}
		return a
	}

	for b != 0 {
		a, b = b, a%b
	}

	if a < 0 {
		return -a
	}

	return a
}

func main() {
	res := GCDEuclidean(48, 18)
	fmt.Println(res)
}
