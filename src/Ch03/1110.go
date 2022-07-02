package main

import (
	"fmt"
)

func main() {
	var n, new, cycle int
	fmt.Scanf("%d", &n)
	cycle = 1
	new = n
	new = (new % 10) * 10 + (((new / 10) + (new % 10)) % 10)

	for new != n {
		new = (new % 10) * 10 + (((new / 10) + (new % 10)) % 10)
		cycle++
	}

	fmt.Println(cycle)
}