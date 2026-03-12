package main

import "fmt"

func lenInt(x int) int {
	count := 0
	for x > 0 {
		x /= 10
		count++
	}
	return count
}

func pow10(n int) int {
	result := 1
	for i := 0; i < n; i++ {
		result *= 10
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func karatsuba(x, y int) int {
	if x < 10 || y < 10 {
		return x * y
	}

	n := max(lenInt(x), lenInt(y))
	m := n / 2

	pow := pow10(m)

	a := x / pow
	b := x % pow
	c := y / pow
	d := y % pow

	z0 := karatsuba(b, d)
	z2 := karatsuba(a, c)
	z1 := karatsuba(a+b, c+d)

	return z2*pow*pow + (z1-z2-z0)*pow + z0
}

func main() {
	a := 1234
	b := 5678

	res := karatsuba(a, b)
	fmt.Println(res)
}
