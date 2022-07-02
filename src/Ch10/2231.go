package main

import "fmt"

func CheckPartial(n int) int {
	check := n
	for n > 0 {
		check += n % 10
		n = n / 10
	}
	return check
}

func main() {
	var N int
	fmt.Scanln(&N)
	for i := 0; i < N; i++ {
		if CheckPartial(i) == N {
			fmt.Println(i)
			return
		}
	}
	fmt.Println(0)
}