package main

import "fmt"

func Fact(N  int) int {
	if N == 0 {
		return 1
	}
	return N * Fact(N - 1)
}

func main() {
	var N int
	fmt.Scanln(&N)
	fmt.Println(Fact(N))
}