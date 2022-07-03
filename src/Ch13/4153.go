package main

import "fmt"

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func IsRight(a, b, c int) {
	hypo := Max(a, Max(b, c))
	X := Min(a, Min(b, c))
	Y := a + b + c - hypo - X
	if ((X * X) + (Y * Y)) == (hypo * hypo) {
		fmt.Println("right")
		return
	}
	fmt.Println("wrong")
}

func main() {
	var a, b, c int
	for {
		fmt.Scan(&a)
		fmt.Scan(&b)
		fmt.Scan(&c)
		if c == 0 {
			break
		}
		IsRight(a, b, c)
	}
}