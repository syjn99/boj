package main

import "fmt"

func main() {
	var x string
	fmt.Scanln(&x)
	fmt.Printf("%d\n", x[0])
}