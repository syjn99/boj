package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var N, x int

	fmt.Fscanln(r, &N)

	A, B := make([]int, N), make([]int, N)

	for i := 0; i < N; i++ {
		fmt.Fscan(r, &x)
		A[i] = x
	}

	for i := 0; i < N; i++ {
		fmt.Fscan(r, &x)
		B[i] = x
	}

	sort.Ints(A)
	sort.Ints(B)

	result := 0

	for i := 0; i < N; i++ {
		result += A[i] * B[N-1-i]
	}

	fmt.Fprintln(w, result)
}