package main

import "fmt"

func main() {
	var a int
	var b int
	var c int

	fmt.Scanf("%d %d %d", &a, &b, &c)
	fmt.Println((a+b)%c)
	fmt.Println( ((a%c) + (b%c))%c  )
	fmt.Println( (a*b)%c ) 
	fmt.Println( ((a%c) * (b%c))%c  )
}