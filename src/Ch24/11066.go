package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func DynPro(dp *[][]int, ind, x int) {
	DP := *dp
	DP[ind][ind] = x
	for i := 0; i < ind; i++ {
		for j := i; j <= ind; j++ {
			DP[ind][i] += DP[j][j]
		}
	}
	for i := ind-1; i >= 0; i-- {
		if i == ind-1 {
			DP[i][ind] = DP[ind][i]
			continue
		}
		min := math.MaxInt
		for j := i; j < ind; j++ {
			// i부터 j까지, j+1부터 ind까지
			cal := DP[j][i] + DP[ind][j+1] // 페이지 수 먼저 더하기
			if i != j {
				cal += DP[i][j]
			}
			if j+1 != ind {
				cal += DP[j+1][ind]
			}
			if cal < min {
				min = cal
			}
		}
		DP[i][ind] = min
	}
	return
}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var T, K, x int

	fmt.Fscanln(r, &T)

	for i := 0; i < T; i++ {
		fmt.Fscanln(r, &K)
		dp := make([][]int, K)
		for j := 0; j < K; j++ {
			dp[j] = make([]int, K)	
		}
		for j := 0; j < K; j++ {
			fmt.Fscan(r, &x)
			DynPro(&dp, j, x)
		}
		fmt.Fscanf(r, "\n")
		fmt.Fprintln(w, dp[0][K - 1])
	}
}