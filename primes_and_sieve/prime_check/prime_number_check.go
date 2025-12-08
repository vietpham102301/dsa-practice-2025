package main

import "fmt"

func isPrime(n int) bool {
	if n < 2 {
		return false
	}

	// giả sử n = p*q
	// mà n = sqrt(n) * sqrt(n)
	// nếu p && q đều lớn hơn sqrt(n) => p*q > n (vô lý)
	// nên nếu n có thể phân tích ra p và q thì một trong hai số phải nhỏ hơn hoặc bằng sqrt(n)
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func main() {
	tests := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	fmt.Printf("result: ")
	for _, t := range tests {
		if isPrime(t) {
			fmt.Printf("%d, ", t)
		}
	}
}
