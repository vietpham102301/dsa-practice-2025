package main

import "fmt"

func isValid(s string) bool {
	var mapping [256]byte
	mapping[')'] = '('
	mapping['}'] = '{'
	mapping[']'] = '['

	stack := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		char := s[i]
		if opener := mapping[s[i]]; opener > 0 {
			if len(stack) == 0 || stack[len(stack)-1] != opener {
				return false
			}

			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, char)
		}
	}

	return len(stack) == 0
}

func main() {
	s := "()[]{}"
	res := isValid(s)
	fmt.Println(res)
}
