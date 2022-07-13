package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var N, x int

	fmt.Fscanln(r, &N)
	Cities, Roads := make([]int, N), make([]int, N)

	for i := 0; i < N - 1; i++ {
		fmt.Fscan(r, &x)
		Roads[i] = x
	}

	Roads = append(Roads, 0)

	for i := 0; i < N; i++ {
		fmt.Fscan(r, &x)
		Cities[i] = x
	}

	fee, index := 0, 0

	for index < N - 1 {
		perLitter := Cities[index]
		kilo := Roads[index]
		for i := index + 1; i < N; i++ {
			if perLitter > Cities[i] {
				index = i
				break
			}
			kilo += Roads[i]
			index = i
		}
		fee += kilo * perLitter
	}

	fmt.Fprintln(w, fee)
}