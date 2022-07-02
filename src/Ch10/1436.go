package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var N int
	fmt.Scanln(&N)
	// arr := make([]int, N + 1)
	// arr[1] = 666
	// for i := 2; i <= N; i++ {
	// 	cand := arr[i - 1] + 1000
	// }
	var cnt int
	for i := 666; i < 10000000; i++ {
		str := strconv.Itoa(i)
		if strings.Contains(str, "666") {
			cnt++
			if cnt == N {
				fmt.Println(i)
				return
			}
		}
	}
}