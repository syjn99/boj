package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	var (
		N, M, cnt int
		s string
	)

	fmt.Fscanf(r, "%d %d\n", &N, &M)

	S := make(map[string]bool)

	for i := 0; i < N; i++ {
		fmt.Fscanln(r, &s)
		S[s] = true
	}

	for i := 0; i < M; i++ {
		fmt.Fscanln(r, &s)
		if S[s] {
			cnt++
		}
	}

	fmt.Fprintln(w, cnt)

	w.Flush()
}