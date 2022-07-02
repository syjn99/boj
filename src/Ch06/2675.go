package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	var (
		T, R int
		S string
	)

	fmt.Fscanln(r, &T)

	for i := 0; i < T; i++ {
		fmt.Fscan(r, &R)
		fmt.Fscan(r, &S)
		for _, c := range S {
			for j := 0; j < R; j++ {
				fmt.Fprintf(w, "%c", c)
			}
			
		}
		fmt.Fprintf(w, "\n")
	}
	w.Flush()
}