package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var N, M, K int

	fmt.Fscanln(r, &N, &M)

	A := make([][]string, N)

	for i := 0; i < N; i++ {
		line, _ := r.ReadString('\n')
		line = strings.TrimSpace(line)
		A[i] = strings.Split(line, " ")
	}

	fmt.Fscanln(r, &M, &K)

	B := make([][]string, M)

	for i := 0; i < M; i++ {
		line, _ := r.ReadString('\n')
		line = strings.TrimSpace(line)
		B[i] = strings.Split(line, " ")
	}

	result := make([][]int, N)
	for i := 0; i < N; i++ {
		result[i] = make([]int, K)
	}

	for i := 0; i < N; i++ {
		row := A[i]
		for j := 0; j < K; j++ {
			elem := 0
			for k := 0; k < M; k++ {
				aNum, _ := strconv.Atoi(row[k])
				bNum, _ := strconv.Atoi(B[k][j])
				elem += aNum * bNum
			}
			result[i][j] = elem
		}
	}

	for _, row := range result {
		for _, elem := range row {
			fmt.Fprintf(w, "%d ", elem)
		}
		fmt.Fprintf(w, "\n")
	}
}