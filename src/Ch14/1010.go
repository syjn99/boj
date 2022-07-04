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

func nCk(n, k int) int {
	if k >= n / 2 {
		k = n - k
	}
	arr := make([]int, k)
	for i := 0; i < k; i++ {
		arr[i] = n
		n--
	}


	for j := k; j >= 1; j-- {
		div := j
		for i := 0; (i < k) && (div != 1); i++ {
			cd := GCD(arr[i], div)
			if cd != 1 {
				arr[i] = arr[i] /  cd
				div = div / cd
			}
		}
	}

	final := 1
	for _, v := range arr {
		final *= v
	}
	return final
}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	var T, N, M int
	fmt.Fscanln(r, &T)
	for i := 0; i < T; i++ {
		fmt.Fscanln(r, &N, &M)
		fmt.Fprintln(w, nCk(M, N))
	}
	w.Flush()
}