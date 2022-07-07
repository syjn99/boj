package main

import (
	"bufio"
	"fmt"
	"os"
)

var Tree [][]int

func Max(arr []int) int {
	max := arr[0]
	for _, v := range arr {
		if v > max {
			max = v
		}
	}
	return max
}

func MaxTwo(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func MaxRoute(n int) int {
	if n == 1 {
		return Tree[0][0]
	}

	DynTree := make([][]int, n)
	DynTree[0] = append(DynTree[0], Tree[0][0])
	for i := 1; i < n; i++ {
		val := Tree[i][0] + DynTree[i - 1][0]
		DynTree[i] = append(DynTree[i], val)
		for j := 1; j < i; j++ {
			val = MaxTwo(DynTree[i - 1][j - 1], DynTree[i - 1][j]) + Tree[i][j]
			DynTree[i] = append(DynTree[i], val)
		}
		val = Tree[i][i] + DynTree[i - 1][i - 1]
		DynTree[i] = append(DynTree[i], val)
	}
	return Max(DynTree[n - 1])
}


func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var n, x int
	fmt.Fscanln(r, &n)
	Tree = make([][]int, n)

	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			fmt.Fscan(r, &x)
			Tree[i] = append(Tree[i], x)
		}
	}

	fmt.Fprintln(w, MaxRoute(n))
}