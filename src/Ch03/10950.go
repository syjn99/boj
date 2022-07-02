package main

import (
	"fmt"
)

func main() {
	var T, A, B int
	fmt.Scanf("%d\n", &T)
	for i := 0; i < T; i++ {
		fmt.Scanf("%d %d\n", &A, &B)
		fmt.Println(A+B)
	}
}