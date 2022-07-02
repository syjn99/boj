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
		lenA, lenB, x, cnt int
	)

	fmt.Fscanf(r, "%d %d\n", &lenA, &lenB)

	mapA, mapB := make(map[int]bool), make(map[int]bool)
	A, B := make([]int, lenA), make([]int, lenB)


	for i := 0; i < lenA; i++ {
		fmt.Fscan(r, &x)
		mapA[x] = true
		A[i] = x
	}

	for i := 0; i < lenB; i++ {
		fmt.Fscan(r, &x)
		mapB[x] = true
		B[i] = x
	}

	for _, elem := range B {
		if !mapA[elem] {
			cnt++
		}
	}

	for _, elem := range A {
		if !mapB[elem] {
			cnt++
		}
	}


	fmt.Fprintln(w, cnt)

	w.Flush()
}