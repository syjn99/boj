package main

import "fmt"

func Min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func Recolor(board *[][]byte, x, y int) int {
	newBoard := *board
	var num1, num2 int
	for i := x; i < x + 8; i++ {
		for j := y; j < y + 8; j++ {
			if ((i + j) % 2 == 0) && (newBoard[i][j] == 66) {
				num1++
			} else if ((i + j) % 2 != 0) && (newBoard[i][j] != 66) {
				num1++
			} else if ((i + j) % 2 == 0) && (newBoard[i][j] != 66) {
				num2++
			} else if ((i + j) % 2 != 0) && (newBoard[i][j] == 66) {
				num2++
			}
		}
	}
	return Min(num1, num2)
}

func main() {
	var (
		M, N int
		s string
	)
	fmt.Scan(&M)
	fmt.Scan(&N)

	board := make([][]byte, M)
	for i := range board {
		board[i] = make([]byte, N)
	}

	for i := 0; i < M; i++ {
		fmt.Scanln(&s)
		for j := 0; j < N; j++ {
			board[i][j] = s[j]
		}
	}

	var minimum int = 2500

	for i := 0; i <= M - 8; i++ {
		for j := 0; j <= N - 8; j++ {
			recoloring := Recolor(&board, i, j)
			if minimum > recoloring {
				minimum = recoloring
			}
		}
	}
	fmt.Println(minimum)
}
