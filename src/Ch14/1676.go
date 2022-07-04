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
	result := 0
	for i := N; i >= 1; i-- {
		num := i
		for {
			if num % 5 == 0 {
				result++
			} else {
				break
			}
			num /= 5
		}
	}
	fmt.Fprintln(w, result)
	w.Flush()
}