package main

import "fmt"

func main() {
	var a int
	var b int
	fmt.Scanf("%d %d", &a, &b)
	if b == 0 {
		return
	}
	fmt.Println(float64(a) / float64(b))
}