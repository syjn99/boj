package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	N, M int
	seqs []string
)

func Sequence2(s string, start, m int, w *bufio.Writer) {
	if m == 1 {
		for i := start; i <= N; i++ {
			if len(s) != 0 {
				fmt.Fprintf(w, "%s %d\n", s, i)
			} else {
				fmt.Fprintf(w, "%d\n", i)
			}
				
		}
		return
	}
	for i := start; i <= N - m + 1; i++ {
		var str string
		if len(s) == 0 {
			str = fmt.Sprintf("%d", i)
		} else {
			str = fmt.Sprintf("%s %d", s, i)
		}
		Sequence2(str, i + 1, m - 1, w)
	}
}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	fmt.Fscanln(r, &N, &M)
	Sequence2("", 1, M, w)
	w.Flush()
}