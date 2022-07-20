package main

import (
	"bufio"
	"fmt"
	"os"
)

const P int = 1000000007

func Power(base, exp, mod int) int {
	if exp == 1 {
		return base % mod
	}
	half := Power(base, exp/2, mod)
	if exp % 2 == 0 {
		return (half * half) % mod
	} else {
		return (((half * half) % mod) * (base % mod)) % mod
	}
}

func nCk(N, K int) int {
	// 분자인 N! / K! = A 에 대해 A mod P를 구하기

	A, B := 1, 1

	for i := N; i > K; i-- {
		A *= i
		A %= P
	}

	// 분모인 (N-K)!에 대해 mod P를 구하기. 페르마의 소정리를 활용, B = (N-K)!라고 한다면 B^-1 mod P = B^(P-2) mod P. 우선 B mod P부터 구하자.

	for i := N - K; i > 0; i-- {
		B *= i
		B %= P
	}

	// B^(P-2)를 구해야 함. 여기서 분할 정복 활용하기
	for i := (P - 2); i > 0; i = i / 2 {
		if i % 2 == 1 {
			A = (A * B) % P
		}
		B = (B * B) % P
	}

	return A
}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var N, K int

	fmt.Fscanln(r, &N, &K)
	fmt.Fprintln(w, nCk(N, K))
}