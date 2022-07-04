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

func GCD(x, y int) int {
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
	return gcd
}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	var N, x int
	fmt.Fscanln(r, &N)
	arr := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(r, &x)
		arr[i] = x
	}
	key := arr[0]
	for i := 1; i < N; i++ {
		gcd := GCD(key, arr[i])
		son, mot := key / gcd, arr[i] / gcd
		fmt.Fprintf(w, "%d/%d\n", son, mot)
	}
	w.Flush()
}