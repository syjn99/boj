package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type Coordinate struct {
	x, y int
}

func Dist(A, B Coordinate) float64 {
	deltaX, deltaY := float64(B.x - A.x), float64(B.y - A.y)
	return math.Sqrt(deltaX * deltaX + deltaY * deltaY)
}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	var (
		T, x1, y1, x2, y2, n, cx, cy, cr, cnt int
	)

	fmt.Fscanln(r, &T)

	for i := 0; i < T; i++ {
		cnt = 0
		fmt.Fscanln(r, &x1,&y1, &x2, &y2)

		A, B := Coordinate{x1, y1}, Coordinate{x2, y2}

		fmt.Fscanln(r, &n)

		for j := 0; j < n; j++ {
			fmt.Fscanln(r, &cx, &cy, &cr)
			C := Coordinate{cx, cy}
			R := float64(cr)
			distAC, distBC := Dist(A, C), Dist(B, C)
			if distAC < R || distBC < R {
				if !((distAC < R) && (distBC < R)) {
					cnt++
				}
			}
			
		}
		fmt.Fprintln(w, cnt)
	}
	w.Flush()
}