package main

import "fmt"

func extendedGCDIterative(a, b int) (int, int, int) {
	r0, r1 := a, b // Tổng quát hoá bài toán thành tìm ri sao a*xi + b*yi = ri; với ri = d
	x0, x1 := 1, 0
	y0, y1 := 0, 1 // a*x0 + b*y0 = a; a*x1 + b*y1 = b

	for r1 != 0 { // dừng tại bước thứ m-1 khi r(m+2) == 0 tức phép chia a % b là phép chia hết
		q := r0 / r1 // chia kiểu int trong Go
		r0, r1 = r1, r0-q*r1
		x0, x1 = x1, x0-q*x1
		y0, y1 = y1, y0-q*y1
	}

	return r0, x0, y0 //g, xi, yi
}

func modInverse(a, m int) (int, error) {
	//normalize a
	a = (a%m + m) % m

	g, x, _ := extendedGCDIterative(a, m)
	if g != 1 {
		return -1, fmt.Errorf("no modular inverse: gcd(%d, %d) = %d != 1", a, m, g)
	}

	x = (x%m + m) % m

	return x, nil
}

func main() {
	a, m := 3, 11
	inv, err := modInverse(a, m)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Inverse of %d mod %d = %d\n", a, m, inv)
	fmt.Printf("Check: (%d * %d) mod %d = %d\n", a, inv, m, (a*inv)%m)
}
