package main

import (
	"fmt"
)

func main() {
	var H, M int
	fmt.Scanf("%d %d", &H, &M)
	if M >= 45 {
		M = M - 45
	} else if H != 0 {
		H = H - 1
		M = M + 15
	} else {
		H = 23
		M = M + 15
	}

	fmt.Printf("%d %d\n", H, M)
}