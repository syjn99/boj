package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var (
		N, p int
		Ps []int
	)
	fmt.Fscanln(r, &N)
	Ps = make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(r, &p)
		Ps[i] = p
	}

	sort.Ints(Ps)

	sum := 0
	Sums := make([]int, N)

	for i := 0; i < N; i++ {
		sum = sum + Ps[i]
		Sums[i] = sum
	}

	result := 0

	for _, v := range Sums {
		result += v
	}

	fmt.Fprintln(w, result)
}