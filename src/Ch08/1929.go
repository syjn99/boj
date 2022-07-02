package main

import (
	"bufio"
	"fmt"
	"os"
)

var primes map[int]int

func main() {
	var M, N int
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	fmt.Fscan(r, &M)
	fmt.Fscan(r, &N)

	primes := make(map[int]int)
	primes[1] = 1

	for j := 2; j * j <= N; j++ {
		if primes[j] == 0 {
			for k := j + j; k <= N; k += j {
				primes[k] = 1
			}
		}
	}

	for i := M; i <= N; i++ {
		if primes[i] == 0 {
			fmt.Fprintln(w, i)
		}
	}

	w.Flush()
}