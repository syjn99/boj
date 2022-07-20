package main

import (
	"bufio"
	"fmt"
	"os"
)

const P int = 1000

func MatMul(A, B [][]int) [][]int {
	N := len(A)
	result := make([][]int, N)
	for i := 0; i < N; i++ {
		result[i] = make([]int, N)
	}

	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			elem := 0
			for k := 0; k < N; k++ {
				a, b := A[i][k], B[k][j]
				elem += a * b
			}
			result[i][j] = elem % P
		}
	}

	return result
}

func MatPower(mat [][]int, B int) [][]int {
	if B == 1 {
		return mat
	}
	half := MatPower(mat, B / 2)
	if B % 2 == 0 {
		return MatMul(half, half)
	} else {
		return MatMul(MatMul(half, half), mat)
	}
}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var N, B, x int

	fmt.Fscanln(r, &N, &B)

	mat := make([][]int, N)
	for i := 0; i < N; i++ {
		mat[i] = make([]int, N)
	}

	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			fmt.Fscan(r, &x)
			mat[i][j] = x
		}
	}

	result := MatPower(mat, B)
	// result := MatMul(mat, mat)

	for _, row := range result {
		for _, elem := range row {
			if B == 1 {
				fmt.Fprintf(w, "%d ", elem % P)
				continue
			}
			fmt.Fprintf(w, "%d ", elem)
		}
		fmt.Fprintf(w, "\n")
	}
}