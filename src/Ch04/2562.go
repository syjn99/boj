package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	var x int

	arr := make([]int, 9, 9)

	for i := 0; i < 9; i++ {
		fmt.Fscanln(r, &x)
		arr[i] = x
	}

	max, index := arr[0], 1

	for i, v := range arr {
		if v > max {
			max = v
			index = i + 1
		}
	}

	fmt.Fprintln(w, max)
	fmt.Fprintln(w, index)
	w.Flush()
}