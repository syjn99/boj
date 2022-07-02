package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func PrintArr(arr []int, w *bufio.Writer) {
	for _, v := range arr {
		fmt.Fprintln(w, v)
	}
}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	var N, x int
	fmt.Fscanln(r, &N)
	arr := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscanln(r, &x)
		arr[i] = x
	}
	
	// SelectionSort(arr)
	sort.Ints(arr)
	PrintArr(arr, w)
	w.Flush()
}