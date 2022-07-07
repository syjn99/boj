package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)


var costs [][]int

func Min(arr3 []int) int {
	return int(math.Min(math.Min(float64(arr3[0]), float64(arr3[1])),float64(arr3[2])))
}

func Coloring(N int) (result int) {
	if N == 1 {
		return Min(costs[0])
	}

	tracking := make([][]int, N)
	tracking[0] = costs[0]
	for i := 1; i < N; i++ {
		tracking[i] = make([]int, 3)
		tracking[i][0] = int(math.Min(float64(tracking[i-1][1]), float64(tracking[i-1][2]))) + costs[i][0]
		tracking[i][1] = int(math.Min(float64(tracking[i-1][0]), float64(tracking[i-1][2]))) + costs[i][1]
		tracking[i][2] = int(math.Min(float64(tracking[i-1][0]), float64(tracking[i-1][1]))) + costs[i][2]
	}

	return Min(tracking[N - 1])
}


func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	var N, R, G, B int
	fmt.Fscanln(r, &N)
	costs = make([][]int, N)

	for i := 0; i < N; i++ {
		fmt.Fscanln(r, &R, &G, &B)
		costs[i] = []int{R, G, B}
	}

	fmt.Fprintln(w, Coloring(N))

	defer w.Flush()
}