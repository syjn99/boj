package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var N, M, x, start, end int
	fmt.Fscanln(r, &N, &M)
	arr := make([]int, N + 1)
	sum := make([]int, N + 1)

	for i := 1; i <= N; i++ {
		fmt.Fscanf(r, "%d", &x)
		arr[i] = x
		sum[i] = sum[i - 1] + x
	}

	for i := 0; i < M; i++ {
		fmt.Fscan(r, &start, &end)
		result := sum[end] - sum[start] + arr[start]
		fmt.Fprintln(w, result)
	}
}