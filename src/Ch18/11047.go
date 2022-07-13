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
	coins := make([]int, N)

	for i := 0; i < N; i++ {
		fmt.Fscanln(r, &x)
		coins[i] = x
	}

	index := N - 1
	cnt := 0

	for K > 0 {
		coin := coins[index]
		if K >= coin {
			K -= coin
			cnt++
			continue
		}
		index--
	}

	fmt.Fprintln(w, cnt)
}