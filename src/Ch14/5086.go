package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	var x, y int
	for {
		fmt.Fscanln(r, &x, &y)
		if x == 0 && y == 0 {
			break
		}
		if y % x == 0 {
			fmt.Fprintln(w, "factor")
		} else if x % y == 0 {
			fmt.Fprintln(w, "multiple")
		} else {
			fmt.Fprintln(w, "neither")
		}
	}
	w.Flush()
}