package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	var X, sum, dist int

	fmt.Fscanln(r, &X)

	for i := 0; ; i++ {
		if X > (i * (i + 1) / 2) && X <= ((i + 1) * (i + 2) / 2 ) {
			dist = X - (i * (i + 1) / 2)
			sum = i + 2
			break
		}
	}

	if sum % 2 == 0 {
		fmt.Fprintf(w, "%d/%d\n", sum - dist, dist)
	} else {
		fmt.Fprintf(w, "%d/%d\n", dist, sum - dist)
	}

	w.Flush()
}