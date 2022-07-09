package main

import (
	"bufio"
	"fmt"
	"os"
)

var arr []int
var LISArr []int = []int{1}

func MaxArr(arr []int) int {
	max := arr[0]
	for _, v := range arr {
		if v > max {
			max = v
		}
	}
	return max
}

func LIS(N int) int {
	for i := 1; i < N; i++ {
		max := 0
		for j := 0; j < i; j++ {
			if (arr[i] > arr[j]) && (LISArr[j] > max) {
				max = LISArr[j]
			}
		}
		LISArr = append(LISArr, 1 + max)
	}
	return MaxArr(LISArr)
}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	
	var N, x int
	fmt.Fscanln(r, &N)
	arr = make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(r, &x)
		arr[i] = x
	}
	fmt.Fprintln(w, LIS(N))
}