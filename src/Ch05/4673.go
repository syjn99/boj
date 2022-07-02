package main

import (
	"bufio"
	"fmt"
	"os"
)

func d(n int) int {
	result := n
	for ; n > 0; n /= 10 {
		result += n % 10
	}
	return result
}

func main() {
	w := bufio.NewWriter(os.Stdout)

	isSelf := make(map[int]int)
	for i := 1; i <= 10000; i++ {
		if isSelf[i] <= 1 {
			fmt.Fprintln(w, i)
			for k := d(i); k <= 10000; k = d(k) {
				isSelf[k] = 2
			}
		}
	}
	w.Flush()
}