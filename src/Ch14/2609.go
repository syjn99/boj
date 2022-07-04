package main

import (
	"bufio"
	"fmt"
	"os"
)

func Min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func isCoprime(x, y int) (int, bool) {
	N := Min(x, y)
	for i := 2; i <= N; i++ {
		if x % i == 0 && y % i == 0 {
			return i, false
		}
	}
	return 0, true
}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	var x, y int
	fmt.Fscanln(r, &x, &y)
	gcd := 1
	for {
		num, ok := isCoprime(x, y)
		if ok == true {
			break
		}
		gcd *= num
		x /= num
		y /= num
	}
	fmt.Fprintln(w, gcd)
	fmt.Fprintln(w, gcd*x*y)
	w.Flush()
}