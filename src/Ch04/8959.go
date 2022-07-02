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
		N, score, bonus int
		test string
	)

	fmt.Fscanln(r, &N)

	for i := 0; i < N; i++ {
		score, bonus  = 0, 0
		fmt.Fscanln(r, &test)
		for _, c := range test {
			if c == 'O' {
				score = score + 1 + bonus
				bonus++
			} else {
				bonus = 0
			}
		}
		fmt.Fprintln(w, score)
	}

	w.Flush()
}