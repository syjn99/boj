package main

import (
	"bufio"
	"fmt"
	"os"
)

func Max(arr []int) (max int) {
	max = arr[0]
	for _, x := range arr {
		if x > max {
			max = x
		}
	}	
	return
}

func Min(arr []int) (min int) {
	min = arr[0]
	for _, x := range arr {
		if x < min {
			min = x
		}
	}	
	return
}

func main() {
	var N, x int
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	fmt.Fscanln(r, &N)
	arr := make([]int, N, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(r, &x)
		arr[i] = x
	}
	fmt.Fprintln(w, Min(arr), Max(arr))
	w.Flush()
}