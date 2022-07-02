package main

import (
	"bufio"
	"fmt"
	"os"
)

func Average(arr []int) (avg float64) {
	avg = 0
	for _, v := range arr {
		avg = avg + float64(v)
	}
	avg = avg/ float64(len(arr))
	return
}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	var (
		N, n, x, cnt int
	)

	fmt.Fscanln(r, &N)

	for i := 0; i < N; i++ {
		cnt = 0

		fmt.Fscan(r, &n)
		arr := make([]int, n, n)

		for j := 0; j < n; j++ {
			fmt.Fscan(r, &x)
			arr[j] = x
		}

		avg := Average(arr)

		for _, v := range arr {
			if float64(v) > avg {
				cnt++
			}
		}
		percent := float64(cnt) / float64(n) * 100.0
		fmt.Fprintf(w, "%.3f%%\n", percent)
	}

	w.Flush()
}