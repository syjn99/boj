package main

import (
	"fmt"
	"sort"
)

func Exchange(arr []int, i, j int) {
	temp := arr[j]
	arr[j] = arr[i]
	arr[i] = temp
	return
}

func SelectionSort(arr []int) {
	for i := 0; i < (len(arr) - 1); i++ {
		key := arr[i]
		var smallest int = i
		for j := i + 1; j < len(arr); j++ {
			if key > arr[j] {
				key = arr[j]
				smallest = j
			}
		}
		Exchange(arr, i, smallest)
	}
}

func PrintArr(arr []int) {
	for _, v := range arr {
		fmt.Println(v)
	}
}

func main() {
	var N, x int
	fmt.Scanln(&N)
	arr := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scanln(&x)
		arr[i] = x
	}
	
	// SelectionSort(arr)
	sort.Ints(arr)
	PrintArr(arr)
}