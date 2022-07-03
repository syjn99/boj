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

func isInside(player, point Coordinate, W, H int) bool {
	if (player.x >= point.x) && (player.x <= (point.x + W)) && (player.y >= point.y) && (player.y <= (point.y + H)) {
		return true
	}
	R := H / 2
	C1, C2 := Coordinate{point.x, point.y + R}, Coordinate{point.x + W, point.y + R}
	if Dist(player, C1) <= float64(R) || Dist(player, C2) <= float64(R) {
		return true
	}
	return false
}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	var (
		W, H, X, Y, P, x, y, cnt int
	)

	fmt.Fscanln(r, &W, &H, &X, &Y, &P)
	point := Coordinate{X, Y}

	for i := 0; i < P; i++ {
		fmt.Fscanln(r, &x, &y)
		player := Coordinate{x, y}
		if isInside(player, point, W, H) {
			cnt++
		}
	}

	fmt.Fprintln(w, cnt)

	w.Flush()
}