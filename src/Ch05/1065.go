package main

import (
	"bufio"
	"fmt"
	"os"
)

func isHansu(x int) bool {
	if x < 100 {
		return true
	}
	if x == 1000 {
		return false
	}
	jari := make([]int, 3, 3)
	for i := 0; i < 3; i++ {
		jari[i] = x % 10
		x /= 10
	}
	if (jari[0] - jari[1]) == (jari[1] - jari[2]) {
		return true
	}
	return false
}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	var N, cnt int
	fmt.Fscanln(r, &N)

	cnt = 0

	for i := 1; i <= N; i++ {
		if isHansu(i) {
			cnt++
		}
	}

	fmt.Fprintln(w, cnt)

	w.Flush()
}