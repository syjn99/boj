package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func Mode(arr []int) int {
	if len(arr) == 1 {
		return arr[0]
	}
	cnt := 1
	cntMax := 0
	m := make(map[int][]int)
	for i := 0; i < (len(arr) - 1); i++ {
		if (arr[i] == arr[i + 1]) {
			cnt++
		} else {
			m[cnt] = append(m[cnt], arr[i])
			if cntMax < cnt {
				cntMax = cnt
			}
			cnt = 1
		}
	}
	if cntMax < cnt {
		cntMax = cnt
	}
	m[cnt] = append(m[cnt], arr[len(arr) - 1])
	if (len(m[cntMax])) == 1 {
		return m[cntMax][0]
	}
	return m[cntMax][1]
}

func main() {
	var N, x int
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	fmt.Fscanln(r, &N)

	sum := 0

	arr := make([]int, N)
	
	for i := 0; i < N; i++ {
		fmt.Fscanln(r, &x)
		arr[i] = x
		sum += x
	}

	sort.Ints(arr)

	average := float64(sum) / float64(N)
	if sum < 0 && ((float64(-sum) / float64(N)) < 0.5) {
		average = 0
	}

	fmt.Fprintf(w, "%.0f\n", average)
	fmt.Fprintln(w, arr[N / 2])
	fmt.Fprintln(w, Mode(arr))
	fmt.Fprintln(w, arr[N - 1] - arr[0])
	w.Flush()	
}