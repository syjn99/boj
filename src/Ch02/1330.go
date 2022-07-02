package main

import (
	"fmt"
)

func main() {
	var a int
	var b int
	fmt.Scanf("%d %d", &a, &b)
	dif := a - b
	if dif > 0 {
		fmt.Println(">")
	} else if dif < 0 {
		fmt.Println("<")
	} else {
		fmt.Println("==")
	}
}