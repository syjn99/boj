package main

import (
	"bufio"
	"fmt"
	"os"
)

var cnt1, cnt2 int

func FibRec(n int) int {
	if n == 1 || n == 2 {
		cnt1++
		return 1
	} else {
		return FibRec(n - 1) + FibRec(n - 2)
	}
}

func FibDyn(n int) int {
	if n <= 2 {
		return 1
	}
	fibs := make([]int, n + 1)
	fibs[1], fibs[2] = 1, 1
	for i := 3; i <= n; i++ {
		cnt2++
		fibs[i] = fibs[i - 1] + fibs[i - 2]
	}
	return fibs[n]
}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var n int
	fmt.Fscanln(r, &n)
	_, _ = FibRec(n), FibDyn(n)
	fmt.Fprintln(w, cnt1, cnt2)
}