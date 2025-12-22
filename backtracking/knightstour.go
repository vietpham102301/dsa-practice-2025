package main

import (
	"fmt"
	"sort"
)

type Move struct {
	x, y  int
	count int
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

func knightsTourOptimized(x, y int, board [][]int, step int) bool {
	n := len(board)
	board[x][y] = step

	if step == n*n {
		return true
	}

	var moves []Move
	for i := 0; i < 8; i++ {
		nx, ny := x+xi[i], y+yi[i]
		if nx >= 0 && ny >= 0 && nx < n && ny < n && board[nx][ny] == 0 {
			// Warnsdorff's Rule: Tính xem từ (nx, ny) đi được bao nhiêu nước nữa
			c := countMoves(nx, ny, board, n)
			moves = append(moves, Move{nx, ny, c})
		}
	}

	// Đi vào nơi "nguy hiểm" (ít lựa chọn) trước
	sort.Slice(moves, func(i, j int) bool {
		return moves[i].count < moves[j].count
	})

	// Thử các nước đi đã sắp xếp
	for _, m := range moves {
		if valid := knightsTourOptimized(m.x, m.y, board, step+1); valid {
			return true
		}
	}

	board[x][y] = 0
	return false
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
	size := 8
	board := make([][]int, size)
	for i := range board {
		board[i] = make([]int, size)
	}

	// Bắt đầu từ ô (0,0) với bước số 1
	done := knightsTourOptimized(0, 0, board, 1)

	if done {
		printChessBoard(board)
	} else {
		fmt.Println("Không tìm thấy giải pháp.")
	}
}
