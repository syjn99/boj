package main

import "fmt"

var dp [][]int
const MOD int = 1000000000

func Answer(arr []int) (sum int) {
	for _, v := range arr {
		sum += v % MOD
	}
	sum = sum % MOD
	return
}

func StairNum(N int) int {
	for i := 2; i <= N; i++ {
		dp[i][0] = dp[i - 1][1]
		dp[i][9] = dp[i - 1][8]
		for j := 1; j <= 8; j++ {
			dp[i][j] = (dp[i - 1][j - 1] + dp[i - 1][j + 1]) % MOD
		}
	}
	return Answer(dp[N])
}

func main() {
	var N int
	fmt.Scan(&N)
	dp = make([][]int, N + 1)
	for i := 0; i < N + 1; i++ {
		dp[i] = make([]int, 10)
	}

	for i := 1; i < 10; i++ {
		dp[1][i] = 1
	}
	fmt.Println(StairNum(N))
}