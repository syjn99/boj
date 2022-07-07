package main

import (
	"bufio"
	"fmt"
	"os"
)

func Pado(N int) int {
	if N <= 5 {
		if N <= 3 {
			return 1
		} else {
			return 2
		}
	}
	cases := make([]int, N + 1)
	for i := 1; i <= 5; i++ {
		cases[i] = Pado(i)
	}
	for i := 6; i <= N; i++ {
		cases[i] = cases[i - 1] + cases[i - 5]
	}
	return cases[N]
}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	
	var T, N int
	fmt.Fscanln(r, &T)
	for i := 0; i < T; i++ {
		fmt.Fscanln(r, &N)
		fmt.Fprintln(w, Pado(N))
	}
}