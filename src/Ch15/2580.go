package main

import (
	"bufio"
	"fmt"
	"os"
)

type Coordinate struct {
	row, col int
}

var board [][]int = make([][]int, 9)
var coordinates []Coordinate

func CheckRow(row, col, j int) bool {
	for i := 0; i < 9; i++ {
		num := board[row][i]
		if num == j {
			return false
		}
	}
	return true
}

func CheckCol(row, col, j int) bool {
	for i := 0; i < 9; i++ {
		num := board[i][col]
		if num == j {
			return false
		}
	}
	return true
}

func CheckBox(row, col, j int) bool {

	rowStart, colStart := (row / 3) * 3, (col / 3) * 3

	for i := rowStart; i < rowStart + 3; i++ {
		for k := colStart; k < colStart + 3; k++ {
			num := board[i][k]
			if num == j {
				return false
			}
		}
	}
	return true
}

func Check(row, col, j int) bool {
	if CheckRow(row, col, j) && CheckCol(row, col, j) && CheckBox(row, col, j) {
		return true
	}
	return false
}

func Sudoku(N int, w *bufio.Writer) {
	if N == len(coordinates) {
		for i := 0; i < 9; i++ {
			for j := 0; j < 9; j++ {
				fmt.Fprintf(w, "%d ", board[i][j])
			}
			fmt.Fprintf(w, "\n")
		}
		w.Flush()
		os.Exit(0)
	}

	row, col := coordinates[N].row, coordinates[N].col
	// if row >= 9 {
	// 	for i := 0; i < 9; i++ {
	// 		for j := 0; j < 9; j++ {
	// 			fmt.Fprintf(w, "%d ", board[i][j])
	// 		}
	// 		fmt.Fprintf(w, "\n")
	// 	}
	// 	w.Flush()
	// 	os.Exit(0)
	// }

	// if col == 9 {
	// 	Sudoku(row + 1, 0, w)
	// }

	// var colIndex int = -1
	// for i := col; i < 9; i++ {
	// 	if board[row][i] == 0 {
	// 		colIndex = i
	// 		break
	// 	}
	// }

	// if colIndex == -1 {
	// 	Sudoku(row + 1, 0, w)
	// }

	// slice := []int{5, 6, 7, 8, 9, 1, 2, 3, 4}

	// // for i := range slice {
  // //   j := rand.Intn(i + 1)
  // //   slice[i], slice[j] = slice[j], slice[i]
	// // }

	for j := 1; j <= 9; j++ {
		if Check(row, col, j) {
			board[row][col] = j
			Sudoku(N + 1, w)
			board[row][col] = 0
		}
	}
}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	var x int

	for i := 0; i < 9; i++ {
		board[i] = make([]int, 9)
	}
	
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fmt.Fscan(r, &x)
			if x == 0 {
				coordinates = append(coordinates, Coordinate{i, j})
			}
			board[i][j] = x
		}
	}
	Sudoku(0, w)
}