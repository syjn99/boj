package main

import (
	"bufio"
	"fmt"
	"os"
)


type Object struct {
	W, V int
}

var (
	N, K int
	objects []Object
	dp [][]int
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

	fmt.Fscanln(r, &N, &K)

	objects = make([]Object, N + 1)
	dp = make([][]int, N + 1)
	for i := 0; i < (N + 1); i++ {
		dp[i] = make([]int, K + 1)
	}

	var W, V int
	
	for i := 1; i <= N; i++ {
		fmt.Fscanln(r, &W, &V)
		obj := Object{W, V}
		objects[i] = obj
	}

	for i := 1; i <= N; i++ { // 물건을 쓸 수 있는...
		obj := objects[i]
		for j := 1; j <= K; j++ { // 무게
			if obj.W > j {
				dp[i][j] = dp[i - 1][j]
			} else {
				dp[i][j] = Max(dp[i - 1][j], dp[i - 1][j - obj.W] + obj.V)
			}
		}
	}
	fmt.Fprintln(w, dp[N][K])

}