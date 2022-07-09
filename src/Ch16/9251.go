package main

import (
	"bufio"
	"fmt"
	"os"
)

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var s1, s2 string

	fmt.Fscanln(r, &s1)
	fmt.Fscanln(r, &s2)

	lenS1, lenS2 := len(s1), len(s2)

	dp := make([][]int, lenS1 + 1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, lenS2 + 1)
	}

	for i := 1; i <= lenS1; i++ {
		cS1 := s1[i - 1]
		for j := 1; j <= lenS2; j++ {
			cS2 := s2[j - 1]
			if cS1 == cS2 {
				dp[i][j] = dp[i - 1][j - 1] + 1
			} else {
				dp[i][j] = Max(dp[i][j - 1], dp[i - 1][j])
			}
		}
	}

	fmt.Fprintln(w, dp[lenS1][lenS2])
}