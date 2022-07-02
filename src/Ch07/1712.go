package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	var A, B, C int

	fmt.Fscan(r, &A)
	fmt.Fscan(r, &B)
	fmt.Fscan(r, &C)

	if B > C {
		fmt.Fprintln(w, -1)
		w.Flush()
		return
	}

	MC := C - B

	if MC == 0 {
		fmt.Fprintln(w, -1)
		w.Flush()
		return
	}

	fmt.Fprintln(w, (A / MC) + 1)
	w.Flush()

}