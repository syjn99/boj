// Bottom-Up
// memo[k] => k원을 만들기 위한 경우의 수.
// 동전의 가치를 작은 수부터 기록하기.

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var (
	coins []int
	memo []int
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var n, k, x int

	fmt.Fscanln(r, &n, &k)

	memo = make([]int, k+1)
	memo[0] = 1
	coins = make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Fscanln(r, &x)
		coins[i] = x
	}

	sort.Ints(coins)

	for i := 0; i < n; i++ {
		coin := coins[i]
		for j := 1; j <= k; j++ {
			if j - coin >= 0 {
				memo[j] += memo[j - coin]
			}
		}
	}
	fmt.Fprintln(w, memo[k])
}