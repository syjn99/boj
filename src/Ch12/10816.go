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

	m := make(map[int]int)

	for i := 0; i < N; i++ {
		fmt.Fscan(r, &x)
		m[x]++
	}


	fmt.Fscan(r, &M)

	for i := 0; i < M; i++ {
		fmt.Fscan(r, &x)
		fmt.Fprintf(w, "%d ", m[x])
	}

	fmt.Fprintf(w, "\n")

	w.Flush()
}