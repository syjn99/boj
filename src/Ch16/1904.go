package main

import (
	"bufio"
	"fmt"
	"os"
)

func Tile(N int) int {
	if N <= 2 {
		return N
	}
	cases := make([]int, N + 1)
	cases[1], cases[2] = 1, 2
	for i := 3; i <= N; i++ {
		cases[i] = (cases[i - 1] + cases[i - 2]) % 15746
	}
	return cases[N]
}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	
	var N int

	fmt.Fscanln(r, &N)
	fmt.Fprintln(w, Tile(N))
}