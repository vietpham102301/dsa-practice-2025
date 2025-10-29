package main

import (
	"fmt"
)

// extendedGCD trả về (gcd, x, y) sao cho a*x + b*y = gcd(a,b)
func extendedGCD(a, b int) (int, int, int) {
	fmt.Printf("→ Gọi extendedGCD(%d, %d)\n", a, b)

	if b == 0 {
		fmt.Printf("  Base case: b = 0 → gcd = %d, x = 1, y = 0\n", a)
		return a, 1, 0
	}

	// Gọi đệ quy với (b, a % b)
	gcd, x1, y1 := extendedGCD(b, a%b)

	// Quay lui — tính x, y mới
	q := a / b
	x := y1
	y := x1 - q*y1

	fmt.Printf("  ← Truy ngược: a=%d, b=%d, q=%d, x=%d, y=%d (vì x=y1, y=x1-q*y1)\n", a, b, q, x, y)
	return gcd, x, y
}

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

func main() {
	a, b := 99, 78
	g, x, y := extendedGCD(a, b)

	fmt.Println("\n=======================")
	fmt.Printf("Kết quả cuối cùng:\n")
	fmt.Printf("gcd(%d, %d) = %d\n", a, b, g)
	fmt.Printf("Hệ số: x = %d, y = %d\n", x, y)
	fmt.Printf("Kiểm tra: %d*%d + %d*%d = %d\n", a, x, b, y, a*x+b*y)
}
