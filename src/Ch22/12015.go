package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var N, x int
	
	fmt.Fscanln(r, &N)
	arr := make([]int, N)

	for i := 0; i < N; i++ {
		fmt.Fscan(r, &x)
		arr[i] = x
	}
	
	lis := make([]int, 0)
	lis = append(lis, 0)
	
	for i := 0; i < N; i++ {
		v, last := arr[i], lis[len(lis) - 1]
		if v > last {
			lis = append(lis, v)
			continue
		}
		left, right := 0, len(lis) - 1
		for {
			mid := (left + right) / 2
			if mid == 0 {
				left = mid+1
				continue
			}
			if v <= lis[mid-1] {
				right = mid-1
				continue
			} else if v > lis[mid] {
				left = mid+1
				continue
			} else {
				lis[mid] = v
				break
			}
		}
	}
	fmt.Fprintln(w, len(lis) - 1)
	
}