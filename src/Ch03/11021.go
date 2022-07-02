package main

import (
	"fmt"
)

func main() {
	var T, A, B int
	fmt.Scanf("%d\n", &T)
	for i := 1; i <= T; i++ {
		fmt.Scanf("%d %d\n", &A, &B)
		fmt.Printf("Case #%d: %d\n", i, A+B)
	}
}