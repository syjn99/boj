package main

import (
	"bufio"
	"fmt"
	"os"
)

func Sequence(w *bufio.Writer, N, M, index int, seq []int) {
	if index == M {
		for i := 1; i <= M; i++ {
			fmt.Fprintf(w, "%d ", seq[i])
		}
		fmt.Fprintf(w, "\n")
		return
	}
	for i := seq[index]; i <= N; i++ {
		if i == 0 {
			continue
		}
		seq[index + 1] = i
		Sequence(w, N, M, index + 1, seq)
	}
}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var N, M int

	fmt.Fscanln(r, &N, &M)

	seq := make([]int, M + 1)

	Sequence(w, N, M, 0, seq)
}