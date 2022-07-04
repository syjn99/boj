// 처음에는 주어진 수를 팩토리얼로 돌면서 2와 5의 개수를 파악하려고 했으나... 그러면 시간 초과가 떠버린다. ㅜㅜ 그래서 tmp += n/5, n/25... 이런식으로 2와 5의 개수를 빠르게 셀 수 있었다.

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

func count(x, n int) (int) {
	num := n
	cnt := 0
	for {
		if x / num == 0 {
			break
		}
		cnt += x / num
		num *= n
	}
	return cnt
}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	var N, M int
	fmt.Fscanln(r, &N, &M)
	if M >= N / 2 {
		M = N - M
	}
	K := N - M
	son2, mot1_2, mot2_2 := count(N, 2), count(M, 2), count(K, 2)
	son5, mot1_5, mot2_5 := count(N, 5), count(M, 5), count(K, 5)
	mot2, mot5 := mot1_2 + mot2_2, mot1_5 + mot2_5
	twos, fives := son2 - mot2, son5 - mot5
	fmt.Println(Min(twos, fives))
	w.Flush()
}