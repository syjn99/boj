package main

import (
	"fmt"
)

func main() {
	var x, y int
	fmt.Scanf("%d\n%d\n", &x, &y)
	if x * y > 0 {
		if x > 0 {
			fmt.Println(1)
		} else {
			fmt.Println(3)
		}
	} else {
		if x > 0 {
			fmt.Println(4)
		} else {
			fmt.Println(2)
		}
	}
}