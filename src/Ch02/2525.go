package main

import (
	"fmt"
)

func main() {
	var H, M, N int
	fmt.Scanf("%d %d\n%d", &H, &M, &N)
	N_H := N / 60
	N_M := N % 60
	H = H + N_H
	M = M + N_M
	if M >= 60 {
		M = M - 60
		H = H + 1
	}
	if H >= 24 {
		H = H - 24
	}
	fmt.Printf("%d %d\n", H, M)
}