package main

import (
	"bufio"
	"fmt"
	"os"
)

var primes map[int]int

func main() {
	var T, n int
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	fmt.Fscanln(r, &T)

	primes := make(map[int]int)
	primes[1] = 1

	for i := 0; i < T; i++ {
		fmt.Fscanln(r, &n)
		for j := 2; j * j <= n; j++ {
			if primes[j] == 0 {
				for k := 2 * j; k <= n; k += j {
					primes[k] = 1
				}
			}
		}
		var a, b int
		if n == 4 {
			a, b = 2, 2
			fmt.Fprintln(w, a, b) 
			continue
		}
		for l := 3; l <= (n / 2); l += 2 {
			if (primes[l] == 0) && (primes[n - l] == 0) {
				a = l
				b = n - l
			}
		}
		fmt.Fprintln(w, a, b) 
	}


	w.Flush()

}