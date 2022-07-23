package main

import (
	"bufio"
	"fmt"
	"os"
)

const P int = 1000000007

func Multiply(mat1, mat2 [][]int) [][]int {
	result := [][]int{[]int{0,0}, []int{0,0}}

	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			temp := 0
			for k := 0; k < 2; k++ {
				temp += mat1[i][k] * mat2[k][j]
			}
			result[i][j] = temp % P
		}
	}

	return result
}

func Power(mat [][]int, n int) [][]int {
	if n == 1 {
		return mat
	}
	half := Power(mat, n/2)

	if n % 2 == 0 {
		return Multiply(half, half)
	} else {
		return Multiply(Multiply(half, half), mat)
	}
}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var n int
	fmt.Fscanln(r, &n)

	mat := make([][]int, 2)
	mat[0] = []int{1, 1}
	mat[1] = []int{1, 0}

	mat = Power(mat, n)

	fmt.Fprintln(w, mat[0][1])

}