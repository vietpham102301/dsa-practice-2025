package main

import "fmt"

func gcdOfStrings(str1 string, str2 string) string {
	l1 := len(str1)
	l2 := len(str2)

	if l1 < l2 {
		return helper(str1, str2, l1)
	}

	return helper(str2, str1, l2)
}

func helper(str1 string, str2 string, l1 int) string {
	f := str2[:l1]
	if str1 == f && checker(str1, str2, l1) {
		return str1
	}

	for i := l1 - 1; i > 0; i-- {
		f = str2[:i]
		t := str1[:i]
		if t == f {
			if checker(t, str2, i) {
				return t
			}
		}
	}

	return ""
}

func checker(str1 string, str2 string, l int) bool {
	for i := 0; i < len(str2) && i+l <= len(str2); {
		if str1 != str2[i:i+l] {
			return false
		}
		i += l
	}
	return (len(str2) % len(str1)) == 0
}

func main() {
	str1 := "ABABAB"
	str2 := "ABAB"

	res := gcdOfStrings(str1, str2)
	fmt.Println(res)
}
