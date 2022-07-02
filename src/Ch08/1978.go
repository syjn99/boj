package main

import "fmt"

func isPrime(x int) int {
	if (x == 2) || (x == 3) {
		return 1
	}
	if (x == 1) || (x % 2 == 0) {
		return 0
	}
	for i := 3; i < x; i += 2 {
		if x % i == 0 {
			return 0
		}
	}
	return 1
}

func main() {
	var N, x, cnt int
	fmt.Scanln(&N)
	for i := 0; i < N; i++ {
		fmt.Scan(&x)
		cnt += isPrime(x)
	}
	fmt.Println(cnt)
}