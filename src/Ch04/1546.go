package main

import (
	"bufio"
	"fmt"
	"os"
)

func Max(arr []float64) (max float64) {
	max = arr[0]

	for _, v := range arr {
		if v > max {
			max = v
		}
	}
	return
}

func Average(arr []float64) (avg float64) {
	avg = 0
	for _, v := range arr {
		avg = avg + v
	}
	avg = avg/ float64(len(arr))
	return
}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	var (
		N int
		x float64
	)
	fmt.Fscanln(r, &N)
	arr := make([]float64, N, N)
	
	for i := 0; i < N; i++ {
		fmt.Fscan(r, &x)
		arr[i] = x
	}

	max := Max(arr)

	for i, v := range arr {
		arr[i] = float64(v) / float64(max) * 100.0 
	}

	fmt.Fprintln(w, Average(arr))
	w.Flush()
}