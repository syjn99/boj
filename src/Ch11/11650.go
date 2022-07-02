package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Coordinate struct {
	x, y int
}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	var N, x, y int
	fmt.Fscanln(r, &N)
	arr := make([]Coordinate, N)
	for i := 0; i < N; i++ {	
		fmt.Fscan(r, &x)
		fmt.Fscan(r, &y)
		arr[i] = Coordinate{x, y}
	}
	sort.Slice(arr, func(i, j int) bool {
		if arr[i].x != arr[j].x {
			return arr[i].x < arr[j].x
		} else {
			return arr[i].y < arr[j].y
		}
	})
	for _, coordinate := range arr {
		fmt.Fprintf(w, "%d %d\n", coordinate.x, coordinate.y)
	}
	w.Flush()
}