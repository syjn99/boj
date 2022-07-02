package main

import (
	"fmt"
)

func main() {
	var N, X, x int
	fmt.Scanf("%d %d", &N, &X)
	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &x)
		if x < X {
			fmt.Printf("%d ", x)
		}
	}
	fmt.Printf("\n")
}