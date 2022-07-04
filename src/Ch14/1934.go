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

func LCM(x, y int) int {
	lcm := 1
	for {
		num, ok := isCoprime(x, y)
		if ok == true {
			break
		}
		lcm *= num
		x /= num
		y /= num
	}
	return lcm*x*y
}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	var T, x, y int
	fmt.Fscanln(r, &T)
	for i := 0; i < T; i++ {
		fmt.Fscanln(r, &x, &y)
		fmt.Fprintln(w, LCM(x, y))
	}
	w.Flush()
}