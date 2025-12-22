package main

import (
	"fmt"
	"sort"
)

// Cấu trúc để lưu nước đi kế tiếp và "trọng số" (số nước đi tiếp theo từ nó)
type Move struct {
	x, y  int
	count int // Số lượng nước đi khả thi từ ô (x, y) này
}

var (
	xi = []int{2, 1, 2, 1, -2, -1, -2, -1}
	yi = []int{1, 2, -1, -2, -1, -2, 1, 2}
)

// Hàm đếm số nước đi khả thi từ vị trí (x, y)
func countMoves(x, y int, board [][]int, n int) int {
	count := 0
	for i := 0; i < 8; i++ {
		nx, ny := x+xi[i], y+yi[i]
		if nx >= 0 && ny >= 0 && nx < n && ny < n && board[nx][ny] == 0 {
			count++
		}
	}
	return count
}

func knightsTourOptimized(x, y int, board [][]int, step int, q *bool) {
	n := len(board)

	// Gán giá trị cho ô hiện tại (Sửa lỗi logic: gán trước khi check full)
	board[x][y] = step

	// Base case: Nếu step == tổng số ô -> Đã xong
	if step == n*n {
		*q = true
		return // Thành công
	}

	// Lấy danh sách các nước đi tiếp theo
	var moves []Move
	for i := 0; i < 8; i++ {
		nx, ny := x+xi[i], y+yi[i]
		if nx >= 0 && ny >= 0 && nx < n && ny < n && board[nx][ny] == 0 {
			// Warnsdorff's Rule: Tính xem từ (nx, ny) đi được bao nhiêu nước nữa
			c := countMoves(nx, ny, board, n)
			moves = append(moves, Move{nx, ny, c})
		}
	}

	// QUAN TRỌNG: Sắp xếp các nước đi theo số lượng nước khả thi tăng dần
	// Đi vào nơi "nguy hiểm" (ít lựa chọn) trước
	sort.Slice(moves, func(i, j int) bool {
		return moves[i].count < moves[j].count
	})

	// Thử các nước đi đã sắp xếp
	for _, m := range moves {
		knightsTourOptimized(m.x, m.y, board, step+1, q)
		if *q {
			return // Nếu tìm thấy đường đi thì thoát luôn, không cần backtrack nữa
		}
	}

	// Backtrack: Nếu đi vào ngõ cụt thì trả lại giá trị 0
	board[x][y] = 0
}

func printChessBoard(board [][]int) {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			fmt.Printf("| %2d ", board[i][j])
		}
		fmt.Println("|")
	}
}

func main() {
	size := 8 // Thử luôn với 8x8
	board := make([][]int, size)
	for i := range board {
		board[i] = make([]int, size)
	}

	done := false
	// Bắt đầu từ ô (0,0) với bước số 1
	knightsTourOptimized(0, 0, board, 1, &done)

	if done {
		printChessBoard(board)
	} else {
		fmt.Println("Không tìm thấy giải pháp.")
	}
}
