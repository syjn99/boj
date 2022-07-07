package main

import (
	"bufio"
	"fmt"
	"os"
)

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func SeqSum(arr []int) int {
	length := len(arr)
	sums := make([]int, length)
	sums[0] = arr[0]
	for i := 1; i < length; i++ {
		sums[i] = Max(sums[i - 1] + arr[i], arr[i])
	}
	realMax := sums[0]
	for _, v := range sums {
		if v > realMax {
			realMax = v
		}
	}
	return realMax
}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	
	var n, x int
	fmt.Fscanln(r, &n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &x)
		arr[i] = x
	}
	fmt.Fprintln(w, SeqSum(arr))
}