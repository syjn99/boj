package main

import (
	"bufio"
	"fmt"
	"os"
)

func nC2(n int) int {
	return n * (n - 1) / 2
}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var N, M, x, cnt int

	fmt.Fscanln(r, &N, &M)

	sum := make([]int, N + 1)
	m := make(map[int]int)

	for i := 1; i <= N; i++ {
		fmt.Fscan(r, &x)
		x = x % M
		temp := (sum[i - 1] + x) % M
		sum[i] = temp
		m[temp]++
	}

	for k, v := range m {
		if k == 0 {
			cnt += v
		}
		cnt += nC2(v)
	}

	fmt.Fprintln(w, cnt)
}