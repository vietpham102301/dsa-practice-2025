package main

func superPow(a int, b []int) int {
	const MODE = 1337
	l := len(b)
	if l == 0 {
		return 1
	}

	last := b[l-1]
	remain := b[:l-1]

	// áp dụng tính chất của luỹ thừa biến đổi biểu thức về dạng: ((a^b1)^10*a^b2)^10*...)a^bn
	// sau đó tính mode từ dưới cùng ra bên ngoài với ((((1^10 * a^b1)^10)*a^b2)^10*...)*a^bn
	// áp dụng công thức a*b mode n = (a mode n * b mode n) mode n
	part1 := fastModulo(superPow(a, remain), 10, MODE)
	part2 := fastModulo(a, last, MODE)

	return part1 * part2 % MODE
}

func fastModulo(a, b, mode int) int {
	a = a % mode
	res := 1

	for b > 0 {
		if b&1 == 1 {
			res = res * a % mode
		}

		a = a * a % mode
		b >>= 1
	}

	return res
}
