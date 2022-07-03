package main

import (
	"fmt"
	"math"
)

func main() {
	var R float64
	fmt.Scanln(&R)
	fmt.Println(R * R * math.Pi)
	fmt.Println(2 * R * R)
}