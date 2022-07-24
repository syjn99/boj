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
	
	var K, N, x int
	fmt.Fscanln(r, &K, &N)

	lans, max := make([]int, K), 0

	for i := 0; i < K; i++ {
		fmt.Fscanln(r, &x)
		lans[i] = x
		if x > max {
			max = x
		}
	}

	left, right, result := 1, max, 0

	for left <= right {
		mid := (left + right) / 2
		total := 0
		for _, v := range lans {
			total += (v / mid)
		}
		if total >= N {
			left = mid + 1
			result = mid
		} else {
			right = mid - 1
		}
	}

	fmt.Fprintln(w, result)
}