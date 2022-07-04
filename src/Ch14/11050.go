package main

import "fmt"

func main() {
	var N, K int
	fmt.Scanln(&N, &K)
	if K >= N / 2 {
		K = N - K
	}
	var result int = 1
	for i := 0; i < K; i++ {
		result *= N
		N--
	}
	for j := K; j >= 1; j-- {
		result /= j
	}
	fmt.Println(result)
}