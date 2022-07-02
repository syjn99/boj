package main

import (
	"bufio"
	"fmt"
	"os"
)

var primes map[int]int

func main() {
	var N int
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	primes := make(map[int]int)
	primes[1] = 1

	for {
		cnt := 0
		fmt.Fscanln(r, &N)
		if N == 0 {
			w.Flush()
			return
		}
		for i := 2; i * i <= 2 * N; i++ {
			if primes[i] == 0 {
				for j := 2 * i; j <= 2 * N; j += i {
					primes[j] = 1
				}
			}
		}
		for k := (N + 1); k <= 2 * N; k++ {
			if primes[k] == 0 {
				cnt++
			}
		}
		fmt.Fprintln(w, cnt)
	}
}