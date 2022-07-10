package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Line struct {
	A, B int
}

var (
	Lines []Line
	LIS []int
)

func MaxArr(arr []int) int {
	max := arr[0]
	for _, v := range arr {
		if v > max {
			max = v
		}
	}
	return max
}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var N, A, B int

	fmt.Fscanln(r, &N)
	Lines = make([]Line, N + 1)

	for i := 1; i <= N; i++ {
		fmt.Fscanln(r, &A, &B)
		line := Line{A, B}
		Lines[i] = line
	}

	sort.Slice(Lines, func(i, j int) bool {
		return Lines[i].A < Lines[j].A 
	})

	LIS = append(LIS, 0)

	for i := 1; i <= N; i++ {
		max := 0
		for j := 0; j < i; j++ {
			if Lines[i].B > Lines[j].B && max < LIS[j] {
				max = LIS[j]
			}
		}
		LIS = append(LIS, 1 + max)
	}
	fmt.Fprintln(w, N - MaxArr(LIS))
}