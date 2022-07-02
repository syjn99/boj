package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	var N int

	fmt.Fscanln(r, &N)

	if N == 1 {
		fmt.Fprintln(w, 1)
		w.Flush()
		return
	}

	for i := 0; i < 1000000000; i++ {
		if N > (3*i*i + 3*i + 1) && N <= (3*i*i + 9*i + 7) {
			fmt.Fprintln(w, i + 2)
			w.Flush()
			return
		}
	}
}