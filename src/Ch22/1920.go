package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func isContain(arr []int, x, start, end int) int {
	if start > end {
		return 0
	}
	mid := (start + end) / 2

	if x == arr[mid] {
		return 1
	} else if x < arr[mid] {
		return isContain(arr, x, start, mid-1)
	} else if x > arr[mid] {
		return isContain(arr, x, mid+1, end)
	}
	return 0
}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var N, M, x int

	fmt.Fscanln(r, &N)

	arr := make([]int, N)

	for i := 0; i < N; i++ {
		fmt.Fscan(r, &x)
		arr[i] = x
	}

	sort.Ints(arr)

	fmt.Fscanln(r)

	fmt.Fscanln(r, &M)
	
	for i := 0; i < M; i++ {
		fmt.Fscan(r, &x)
		fmt.Fprintln(w, isContain(arr, x, 0, N-1))
	}
}