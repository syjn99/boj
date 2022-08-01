package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	choos, checks []int
	dp [][]bool
	M int
)

func Fill(weight, ind int) {
	if (ind >= M+1) || (dp[weight][ind] == true) {
		return
	}
	choo := 0
	if ind < M {
		choo = choos[ind+1]
	}
	dp[weight][ind] = true
	Fill(weight+choo, ind+1)
	if weight >= choo {
		Fill(weight-choo, ind+1)
	} else {
		Fill(choo-weight, ind+1)
	}
	Fill(weight, ind+1)
}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	// M은 추의 개수, N은 확인할 구슬들의 개수
	var N, x, totalSum int

	choos = append(choos, 0)

	fmt.Fscanln(r, &M)

	for i := 0; i < M; i++ {
		fmt.Fscan(r, &x)
		totalSum += x
		choos = append(choos, x)
	}
	fmt.Fscanln(r)

	fmt.Fscanln(r, &N)

	resultStr := ""
	maxCheck := 0

	for i := 0; i < N; i++ {
		fmt.Fscan(r, &x)
		if maxCheck < x {
			maxCheck = x
		}
		checks = append(checks, x)
	}
	fmt.Fscanln(r)

	if maxCheck < totalSum {
		maxCheck = totalSum
	}

	dp = make([][]bool, maxCheck+1)
	for i := 0; i < (maxCheck+1); i++ {
		dp[i] = make([]bool, M+1)
	}

	Fill(0, 0)

	for _, v := range checks {
		if dp[v][M] == true {
			resultStr += "Y "
		} else {
			resultStr += "N "
		}
	}

	// for i, v := range dp {
	// 	fmt.Printf("%d번째 줄", i)
	// 	fmt.Println(v)
	// }

	fmt.Fprintln(w, resultStr)
}