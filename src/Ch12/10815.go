package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	var N, M, x int

	fmt.Fscanln(r, &N)

	m := make(map[int]bool)

	for i := 0; i < N; i++ {
		fmt.Fscan(r, &x)
		m[x] = true
	}


	fmt.Fscan(r, &M)

	for i := 0; i < M; i++ {
		fmt.Fscan(r, &x)
		if m[x] {
			fmt.Fprintf(w, "%d ", 1)
		} else {
			fmt.Fprintf(w, "%d ", 0)
		}
	}

	fmt.Fprintf(w, "\n")

	w.Flush()
}