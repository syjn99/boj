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

	var (
		N int
		x string
	)

	fmt.Fscanln(r, &N)
	arr := make([]string, N)

	for i := 0; i < N; i++ {
		fmt.Fscanln(r, &x)
		arr[i] = x
	}

	sort.Slice(arr, func(i, j int) bool {
		a, b := arr[i], arr[j]
		if len(a) != len(b) {
			return len(a) < len(b)
		} else {
			for i := 0; i < len(a); i++ {
				if a[i] == b[i] {
					continue
				} else {
					return a[i] < b[i]
				}
			}
		}
		return true
	})
	
	for i := 0; i < N - 1; i++ {
		if arr[i] == arr[i + 1] {
			continue
		}
		fmt.Fprintln(w, arr[i])
	}
	fmt.Fprintln(w, arr[N-1])
	w.Flush()
}