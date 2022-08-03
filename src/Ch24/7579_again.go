// Bottom-Up

package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	Apps, Costs, CostSum []int
)

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	
	var N, M, x int

	fmt.Fscanln(r, &N, &M)

	Apps, Costs, CostSum = make([]int, N+1), make([]int, N+1), make([]int, N+1)

	sum := 0

	for i := 1; i <= N; i++ {
		fmt.Fscan(r, &x)
		Apps[i] = x
	}

	for i := 1; i <= N; i++ {
		fmt.Fscan(r, &x)
		sum += x
		Costs[i] = x
		CostSum[i] = sum
	}

	dp := make([]int, sum+1)

	for i := 1; i <= N; i++ {
		upperBound := CostSum[i]
		cost := Costs[i]
		for j := (upperBound - cost) ; j >= 0; j-- {
			dp[j + cost] = Max(dp[j + cost], dp[j] + Apps[i])
		}
	}

	answer := 0
	for i := 0; i < sum+1; i++ {
		if dp[i] >= M {
			answer = i
			break
		}
	}

	fmt.Fprintln(w, answer)
}