package main

import "fmt"

func Calc(k, n int) int {
	if k == 0 {
		return n
	} else if n == 1 {
		return 1
	} else {
		return Calc(k - 1, n) + Calc(k, n - 1)
	}
}

func main() {
	var T, k, n int
	fmt.Scanln(&T)

	for i := 0; i < T; i++ {
		fmt.Scanln(&k)
		fmt.Scanln(&n)
		fmt.Println(Calc(k, n))
	}
}