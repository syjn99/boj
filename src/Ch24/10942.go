package main

import (
	"bufio"
	"fmt"
	"os"
)

func Fill(dp *[][]int, nums []int, start, end int) int {
	if start > end {
		return 1
	}

	DP := *dp

	if DP[start][end] != 0 {
		return DP[start][end]
	}

	if start == end {
		DP[start][end] = 1
		return 1
	}

	if nums[start] != nums[end] {
		DP[start][end] = 2
		return 2
	} else {
		result := Fill(dp, nums, start+1, end-1)
		if result == 1 {
			DP[start][end] = 1
			return 1
		} else {
			DP[start][end] = 2
			return 2
		}
	}
}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var N, M, x, start, end int

	fmt.Fscanln(r, &N)
	nums := make([]int, N+1)
	// dp의 초기값 0, 팰린드롬이면 1, 아니면 2로 저장할 예정.
	dp := make([][]int, N+1)
	for i := 0; i < N+1; i++ {
		dp[i] = make([]int, N+1)
	}

	for i := 1; i <= N; i++ {
		fmt.Fscan(r, &x)
		nums[i] = x
	}
	fmt.Fscanln(r)

	fmt.Fscanln(r, &M)
	for i := 0; i < M; i++ {
		fmt.Fscanln(r, &start, &end)
		result := Fill(&dp, nums, start, end)
		if result == 1 {
			fmt.Fprintln(w, 1)
		} else {
			fmt.Fprintln(w, 0)
		}
	}
}