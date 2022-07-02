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
		N, M, cnt int
		name string
	)

	fmt.Fscanf(r, "%d %d\n", &N, &M)

	m := make(map[string]bool)
	var names []string

	for i := 0; i < N; i++ {
		fmt.Fscanln(r, &name)
		m[name] = true
	}

	for i := 0; i < M; i++ {
		fmt.Fscanln(r, &name)
		if m[name] {
			cnt++
			names = append(names, name)
		}
	}

	sort.Strings(names)

	fmt.Fprintln(w, cnt)
	for _, n := range names {
		fmt.Fprintln(w, n)
	}

	w.Flush()
}