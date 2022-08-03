package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	coins []int
	dp [][]int
)

// Top-Down

// D(n, k) : 1~n번까지의 동전으로 k원을 만들 수 있는 경우의 수
// D(n, k) = D(n-1, k) => n번째 동전을 사용하지 않는 경우
//					+ D(n, k - Cn) => n번째 동전을 사용하는 경우. 

// Bottom-Up
// memo[k] => k원을 만들기 위한 경우의 수.
// 동전의 가치를 작은 수부터 기록하기.

func DP(n, k int) int {
	if k < 0 {
		return 0
	}
	
	if n == 0 {
		return 0
	}

	if k == 0 {
		dp[n][0] = 1
		return 1
	}

	if dp[n][k] != 0 {
		return dp[n][k]
	}
	dp[n][k] = DP(n-1, k) + DP(n, k - coins[n])
	return dp[n][k]
}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var n, k, x int

	fmt.Fscanln(r, &n, &k)

	dp = make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, k+1)
	}

	coins = append(coins, 0)

	for i := 0; i < n; i++ {
		fmt.Fscanln(r, &x)
		coins = append(coins, x)
	}

	fmt.Fprintln(w, DP(n, k))

	// for _, v := range dp {
	// 	fmt.Println(v)
	// }
}
