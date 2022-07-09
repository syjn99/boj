package main

import (
	"bufio"
	"fmt"
	"os"
)

var stairs []int
var dp [][]int

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func MaxScore(N int) int {
	for i := 1; i <= N; i++ {
		dp[i][0] = Max(dp[i - 1][1], dp[i - 1][2])
		dp[i][1] = dp[i - 1][0] + stairs[i]
		dp[i][2] = dp[i - 1][1] + stairs[i]
	}

	return Max(dp[N][1], dp[N][2])
}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var N, x int

	fmt.Fscanln(r, &N)

	stairs = make([]int, N + 1)

	dp = make([][]int, N + 1)

	for i := 0; i <= N; i++ {
		dp[i] = make([]int, 3)
	}

	for i := 1; i <= N; i++ {
		fmt.Fscanln(r, &x)
		stairs[i] = x
	}

	fmt.Fprintln(w, MaxScore(N))
}