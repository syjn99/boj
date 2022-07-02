package main

import (
	"bufio"
	"fmt"
	"os"
)

// var primes []int

// func isPrime(x int) int {
// 	if (x == 2) || (x == 3) {
// 		primes = append(primes, x)
// 		return 1
// 	}
// 	if (x == 1) || (x % 2 == 0) {
// 		return 0
// 	}
// 	for i := 3; float64(i) < math.Sqrt(float64(x)); i += 2 {
// 		if x % i == 0 {
// 			return 0
// 		}
// 	}
// 	primes = append(primes, x)
// 	return 1
// }

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	var N int
	fmt.Fscanln(r, &N)
	if N == 1 {
		return
	}
	for N % 2 == 0 {
		N /= 2
		fmt.Fprintln(w, 2)
	}
	for i := 3; i <= N; i += 2 {
		if N == 1 {
			return
		}
		// if isPrime(i) == 0 {
		// 	continue
		// }
		for N % i == 0 {
			N /= i
			fmt.Fprintln(w, i)
		}
	}
	w.Flush()
}