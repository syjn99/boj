package main

import "fmt"

func Fib(N  int) int {
	if N == 0 {
		return 0
	} else if N == 1{
		return 1
	}
	return Fib(N - 1) + Fib(N - 2)
}

func main() {
	var N int
	fmt.Scanln(&N)
	fmt.Println(Fib(N))
}