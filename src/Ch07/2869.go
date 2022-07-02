package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	var A, B, V int
	fmt.Fscan(r, &A)
	fmt.Fscan(r, &B)
	fmt.Fscan(r, &V)

	perDay := A - B
	target := V - A

	days := target / perDay
	if target % perDay != 0 {
		days++
	}

	fmt.Fprintln(w, days + 1)

	w.Flush()
}