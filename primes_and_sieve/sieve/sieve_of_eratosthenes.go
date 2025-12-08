package main

import "fmt"

func sieveOfEratosthenes(n int) []bool {
	isPrime := make([]bool, n+1)
	for i := range isPrime {
		isPrime[i] = true
	}

	isPrime[0], isPrime[1] = false, false

	// tìm các bội của p bắt đầu từ p*p vì các bội số của p với q < p thì đã bị vòng trước đánh dấu (1...q * p)
	// nếu p*p > n tức là đã xét hết bội số q của p (q < p) ở các vòng trước
	// và các bội lớn hơn n thì không cần xét vì vượt ngoài phạm vi xét là n
	// tại p*p = n thì cần đánh dấu số cuối cùng (n) không phải số nguyên tố
	for p := 2; p*p <= n; p++ {
		if isPrime[p] {
			// mark all multiple of p*p as false
			for j := p * p; j <= n; j += p {
				isPrime[j] = false
			}
		}
	}

	return isPrime
}

func main() {
	n := 15

	isPrime := sieveOfEratosthenes(n)
	fmt.Printf("result: ")
	for i := 0; i <= n; i++ {
		if isPrime[i] {
			fmt.Printf("%d, ", i)
		}
	}
}
