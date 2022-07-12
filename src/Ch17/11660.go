package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	matrix [][]int
	sum []int
)

func Convert(x, y int) int {
	return (x - 1) * len(matrix) + y
}

func PartialSum(x, y int) {
	start, end := 0, Convert(x, y)

	for i := 1; i <= end; i++ {
		if sum[i] == 0 {
			start = i
			break	
		}
	}


	N := len(matrix)

	for i := start; i <= end; i++ {
		row := i / N + 1
		col := i % N
		if col == 0 {
			row--
			col = N
		}
		sum[i] = sum[i - 1] + matrix[row - 1][col - 1]
	}
}

func Calculate(x1, y1, x2, y2 int) int {
	result := 0
	if sum[Convert(x2, y2)] == 0 {
		PartialSum(x2, y2)
	}
	for i := x1; i <= x2; i++ {
		start, end := Convert(i, y1), Convert(i, y2)
		temp := sum[end] - sum[start - 1]
		result += temp
	}
	return result
}


func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var N, M, x int
	
	fmt.Fscanln(r, &N, &M)

	matrix = make([][]int, N)
	sum = make([]int, N*N + 1)
	sum[0] = 0

	for i := 0; i < N; i++ {
		matrix[i] = make([]int, N)
	}

	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			fmt.Fscan(r, &x)
			matrix[i][j] = x
		}
	}

	for i := 0; i < M; i++ {
		var x1, y1, x2, y2 int
		fmt.Fscan(r, &x1, &y1, &x2, &y2)
		fmt.Fprintln(w, Calculate(x1, y1, x2, y2))
	}
}