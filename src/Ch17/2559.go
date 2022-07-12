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

	var N, K, x int

	fmt.Fscanln(r, &N, &K)
	arr := make([]int, N + 1)
	sum := make([]int, N + 1)

	for i := 1; i <= N; i++ {
		fmt.Fscan(r, &x)
		arr[i] = x
		sum[i] = sum[i - 1] + x
	}

	max := sum[K] - sum[0]

	for i := K; i <= N; i++ {
		temp := sum[i] - sum[i - K]
		if temp > max {
			max = temp
		}
	}

	fmt.Fprintln(w, max)
}