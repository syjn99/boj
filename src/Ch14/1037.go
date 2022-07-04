package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	var N, x int
	fmt.Fscanln(r, &N)
	arr := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(r, &x)
		arr[i] = x
	}
	sort.Ints(arr)
	fmt.Fprintln(w, arr[0] * arr[N - 1])
	w.Flush()
}