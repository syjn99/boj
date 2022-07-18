package main

import (
	"bufio"
	"fmt"
	"os"
)

func Multiply(A, B, C int) int {
	if B == 1 {
		return A % C
	}

	divide := Multiply(A, B/2, C)
	if B % 2 == 0 {
		return (divide * divide) % C
	} else {
		return ((A % C) * ((divide * divide) % C)) % C
	}
}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var A, B, C int

	fmt.Fscanln(r, &A, &B, &C)

	fmt.Fprintln(w, Multiply(A, B, C))
}