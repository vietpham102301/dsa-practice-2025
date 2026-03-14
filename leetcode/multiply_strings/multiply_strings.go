package main

import "fmt"

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	n1, n2 := len(num1), len(num2)
	res := make([]int, n1+n2)
	for i := n1 - 1; i >= 0; i-- {
		for j := n2 - 1; j >= 0; j-- {
			mul := int(num1[i]-'0') * int(num2[j]-'0')
			p1, p2 := i+j, i+j+1
			sum := mul + res[p2]

			res[p2] = sum % 10
			res[p1] += sum / 10
		}
	}

	ans := ""
	for _, num := range res {
		if len(ans) == 0 && num == 0 {
			continue
		}
		ans += string(rune(num + '0'))
	}

	return ans
}

func main() {
	num1 := "123"
	num2 := "456"

	res := multiply(num1, num2)
	fmt.Println(res)
}
