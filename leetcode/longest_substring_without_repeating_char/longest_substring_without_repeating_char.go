package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	n := len(s)
	if n <= 1 {
		return n
	}

	// Sử dụng mảng 128 phần tử cho toàn bộ bảng mã ASCII
	// Giá trị lưu trữ là: vị trí cũ + 1 (để tận dụng giá trị mặc định là 0)
	var lastIndex [128]int

	maxLength := 0
	left := 0

	for right := 0; right < n; right++ {
		char := s[right] // Lấy byte trực tiếp, không cần decode rune

		// Nếu ký tự này đã xuất hiện bên trong cửa sổ hiện tại
		if lastIndex[char] > left {
			left = lastIndex[char]
		}

		// Cập nhật maxLength
		if right-left+1 > maxLength {
			maxLength = right - left + 1
		}

		// Lưu vị trí hiện tại + 1
		lastIndex[char] = right + 1
	}

	return maxLength
}

func main() {
	s := "abcabcbb"
	res := lengthOfLongestSubstring(s)
	fmt.Println(res)
}
