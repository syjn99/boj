package main

import (
	"bufio"
	"fmt"
	"os"
)

var MemoZero, MemoOne []int

func Count(N int) (zero, one int) {
	if N < 2 {
		return MemoZero[N], MemoOne[N]
	}

	if MemoZero[N] != 0 {
		return MemoZero[N], MemoOne[N]
	}

	firstZero, firstOne := Count(N - 1)
	secondZero, secondOne := Count(N - 2)
	MemoZero[N] = firstZero + secondZero
	MemoOne[N] = firstOne + secondOne

	return Count(N)

}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	MemoZero, MemoOne = make([]int, 41), make([]int, 41)
	MemoZero[0], MemoZero[1] = 1, 0
	MemoOne[0], MemoOne[1] = 0, 1

	var T, N int

	fmt.Fscanln(r, &T)

	for i := 0; i < T; i++ {
		fmt.Fscanln(r, &N)
		a, b := Count(N)
		fmt.Fprintln(w, a, b)
	}
}