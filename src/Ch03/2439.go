package main

import (
	"fmt"
)

func main() {
	var x int
	fmt.Scanf("%d", &x)
	for i := 1; i <= x; i++ {
		for k := 0; k < (x - i); k++ {
			fmt.Printf(" ")
		}
		for j := 0; j < i; j++ {
			fmt.Printf("*")
		}
		fmt.Printf("\n")
	}
}