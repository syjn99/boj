package main

import (
	"fmt"
	"sort"
)

func main() {
	var N string
	fmt.Scanln(&N)
	arr := make([]int, len(N))
	for i, c := range N {
		arr[i] = int(c)
	}
	sort.Ints(arr)
	for i := (len(arr) - 1); i >= 0; i-- {
		fmt.Printf("%c", byte(arr[i]))
	}
	fmt.Printf("\n")
}