package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	arr []int
	bitonic []int
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

func countBitonic(index int) {
	firstSlice := arr[:(index + 1)]
	secondSlice := arr[index:]


	// firstSlice에서 LIS 구하기
	var LIS, LDS []int

	LIS = append(LIS, 1)
	LDS = append(LDS, 1)

	for i := 1; i < len(firstSlice); i++ {
		max := 0
		for j := 0; j < i; j++ {
			if (firstSlice[i] > firstSlice[j]) && (LIS[j] > max) {
				max = LIS[j]
			}
		}
		LIS = append(LIS, 1 + max)
	}

	// secondSlice에서 LDS 구하기
	for i := 1; i < len(secondSlice); i++ {
		max := 0
		for j := 0; j < i; j++ {
			if (secondSlice[i] < secondSlice[j]) && (LDS[j] > max) {
				max = LDS[j]
			}
		}
		LDS = append(LDS, 1 + max)
	}

	IS, DS := MaxArr(LIS), MaxArr(LDS)

	bitonic[index] = IS + DS - 1
}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var N, x int
	
	fmt.Fscanln(r, &N)
	arr = make([]int, N)
	bitonic = make([]int, N)

	for i := 0; i < N; i++ {
		fmt.Fscan(r, &x)
		arr[i] = x	
	}

	for i := 0; i < N; i++ {
		countBitonic(i)
	}

	fmt.Fprintln(w, MaxArr(bitonic))

}