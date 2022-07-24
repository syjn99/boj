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
	
	var N, C, x int
	fmt.Fscanln(r, &N, &C)


	houses := make([]int, N)

	for i := 0; i < N; i++ {
		fmt.Fscanln(r, &x)
		houses[i] = x
	}

	sort.Ints(houses)



	left, right, result := 1, (houses[N-1] - houses[0]), 0

	for left <= right {
		mid := (left + right) / 2 // mid는 거리.
		prev, cnt := houses[0], 1
		for i := 1; i < N; i++ {
			if houses[i] - prev >= mid {
				prev = houses[i]
				cnt++
			}
		}

		if cnt >= C {
			left = mid + 1
			result = mid
		} else {
			right = mid - 1
		}
	}

	fmt.Fprintln(w, result)
}