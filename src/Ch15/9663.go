package main

import "fmt"

var TotalCnt int = 0

func Check(board [][]int, N, row, col int) bool {
	for i := 0; i < row; i++ {
		if board[i][col] == 1 {
			return false
		}
	}
	for i, j := row, col; i >= 0 && j >= 0; i, j = i - 1, j - 1 {
		if board[i][j] == 1 {
			return false
		}
	}
	for i, j := row, col; i >= 0 && j < N; i, j = i - 1, j + 1 {
		if board[i][j] == 1 {
			return false
		}
	}

	return true
}

func Nqueen(board [][]int, N, row int) {
	if row == N {
		TotalCnt++
		return
	}
	for i := 0; i < N; i++ {
		if Check(board, N, row, i) {
			board[row][i] = 1
			Nqueen(board, N, row + 1)
			board[row][i] = 0
		}
	}
}

func main() {
	var N int
	fmt.Scanln(&N)
	board := make([][]int, N)
	for i := 0; i < N; i++ {
		board[i] = make([]int, N)
	}
	Nqueen(board, N, 0)
	fmt.Println(TotalCnt)
}