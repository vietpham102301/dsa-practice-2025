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

func main() {
	a, b := 99, 78
	g, x, y := extendedGCD(a, b)

	fmt.Println("\n=======================")
	fmt.Printf("Kết quả cuối cùng:\n")
	fmt.Printf("gcd(%d, %d) = %d\n", a, b, g)
	fmt.Printf("Hệ số: x = %d, y = %d\n", x, y)
	fmt.Printf("Kiểm tra: %d*%d + %d*%d = %d\n", a, x, b, y, a*x+b*y)
}
