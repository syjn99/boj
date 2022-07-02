package main

import (
	"bufio"
	"fmt"
	"os"
)

func Hanoi(N int) int {
	if N == 1 {
		return 1
	}
	return 2 * Hanoi(N - 1) + 1
}

func PrintProcess(N, src, dst, lay int, w *bufio.Writer) {
	if N == 1 {
		fmt.Fprintln(w, src, dst)
		return
	}
	PrintProcess(N - 1, src, lay, dst, w)
	fmt.Fprintln(w, src, dst)
	PrintProcess(N - 1, lay, dst, src, w)
}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	var N int
	fmt.Fscanln(r, &N)
	fmt.Fprintln(w, Hanoi(N))
	PrintProcess(N, 1, 3, 2, w)
	w.Flush()
}