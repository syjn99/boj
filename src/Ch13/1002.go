package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	
	var T, x1, y1, r1, x2, y2, r2 int

	fmt.Fscanln(r, &T)

	for i := 0; i < T; i++ {
		fmt.Fscanln(r, &x1, &y1, &r1, &x2, &y2, &r2)
		if x1 == x2 && y1 == y2 && r1 == r2 {
			fmt.Fprintln(w, -1)
			continue
		}
		deltaX, deltaY := float64(x2 - x1), float64(y2 - y1)
		d := math.Sqrt(deltaX * deltaX + deltaY * deltaY)
		rSum := float64(r1 + r2)
		rSub := math.Abs(float64(r2 - r1))
		var result int
		if d > rSum {
			result = 0
		} else if d == rSum {
			result = 1
		} else if d < rSum && d > rSub {
			result = 2
		} else if d == rSub {
			result = 1
		} else if d < rSub {
			result = 0
		}
		fmt.Fprintln(w, result)
	}

	w.Flush()
}