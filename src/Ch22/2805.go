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
	
	var N, M, x int
	fmt.Fscanln(r, &N, &M)

	trees, max := make([]int, N), 0

	for i := 0; i < N; i++ {
		fmt.Fscan(r, &x)
		trees[i] = x
		if x > max {
			max = x
		}
	}

	left, right, result := 0, max, 0

	for left <= right {
		mid := (left + right) / 2
		total := 0
		for _, v := range trees {
			if v <= mid {
				continue
			}
			total += (v - mid)
		}
		if total >= M {
			left = mid + 1
			result = mid
		} else {
			right = mid - 1
		}
	}

	fmt.Fprintln(w, result)
}