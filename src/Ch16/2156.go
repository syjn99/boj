package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	podos []int
	dp [][]int
)

func MaxArr(arr []int) int {
	max := arr[0]
	for _, v := range arr {
		if v > max {
			max = v
		}
	}
	return max
}

func MaxDrink(N int) int {
	for i := 1; i <= N; i++ {
		dp[i][0] = MaxArr([]int{dp[i - 1][1], dp[i - 1][2]})
		dp[i][1] = dp[i - 1][0] + podos[i]
		dp[i][2] = MaxArr([]int{dp[i - 1][2], dp[i - 1][1] + podos[i], dp[i - 1][0] + podos[i]})
	}

	return MaxArr(dp[N])
}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var N, x int
	fmt.Fscanln(r, &N)

	podos = make([]int, N + 1)
	dp = make([][]int, N + 1)
	for i := 0; i <= N; i++ {
		dp[i] = make([]int, 3)
	}

	for i := 1; i <= N; i++ {
		fmt.Fscanln(r, &x)
		podos[i] = x
	}


	fmt.Fprintln(w, MaxDrink(N))
}