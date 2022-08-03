package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var (
	dp [][]int
	Apps, Costs, MemorySum []int
)

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Solve(n, k int) int {
	if n < 0 {
		return 0
	}
	if MemorySum[n] < k {
		dp[n][k] = math.MaxInt
		return math.MaxInt
	}
	if k < 0 {
		return 0
	}
	if dp[n][k] != 0 {
		return dp[n][k]
	}
	dp[n][k] = Min(Solve(n-1, k), Solve(n-1, k-Apps[n]) + Costs[n])
	return dp[n][k]
}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var N, M, x int

	fmt.Fscanln(r, &N, &M)

	dp = make([][]int, N+1)
	for i := 0; i < N+1; i++ {
		dp[i] = make([]int, M+1)
	}

	Apps, Costs, MemorySum = make([]int, N+1), make([]int, N+1), make([]int, N+1)

	sum := 0

	for i := 1; i <= N; i++ {
		fmt.Fscan(r, &x)
		sum += x
		Apps[i] = x
		MemorySum[i] = sum
	}

	for i := 1; i <= N; i++ {
		fmt.Fscan(r, &x)
		Costs[i] = x
	}

	fmt.Fprintln(w, Solve(N, M))
}