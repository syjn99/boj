package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scanf("%d\n", &n)
	sum := 0
	for i := 1; i <= n; i++ {
		sum = sum + i
	}
	fmt.Println(sum)
}