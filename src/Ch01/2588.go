package main

import "fmt"

func main() {
	var a int
	var b int

	fmt.Scanf("%d\n%d\n", &a, &b)

	first := a * (b % 10)
	second := a * ( ((b % 100) - (b % 10)) / 10 )
	fourth := a*b
	third := fourth - first - (second * 10)
	third = third / 100
	fmt.Println(first)
	fmt.Println(second)
	fmt.Println(third)
	fmt.Println(fourth)
}