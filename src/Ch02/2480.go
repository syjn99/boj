package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c int
	fmt.Scanf("%d %d %d", &a, &b, &c)
	var reward int
	if a == b {
		if b == c {
			reward = 10000 + a * 1000
		} else {
			reward = 1000 + a * 100
		}
	} else if b == c {
		reward = 1000 + b * 100
	} else if a == c {
		reward = 1000 + a * 100
	} else {
		reward = (int)(math.Max(math.Max(float64(a), float64(b)), float64(c)) * 100)
	}

	fmt.Println(reward)
}