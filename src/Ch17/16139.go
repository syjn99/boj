package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	S string
	alphaToSums map[string][]int
)

func PartialSum(alpha string, end int) []int {
	sum := alphaToSums[alpha]
	length := len(sum)
	if length == 0 {
		sum = append(sum, 0)
		length++
	}
	for i := length; i <= end + 1; i++ {
		if alpha == string(S[i - 1]) {
			sum = append(sum, sum[i - 1] + 1)
		} else {
			sum = append(sum, sum[i - 1])
		}
	}
	alphaToSums[alpha] = sum
	return sum
}

func Calculate(alpha string, start, end int) int {
	sum := alphaToSums[alpha]
	if len(sum) < end + 2 {
		sum = PartialSum(alpha, end)
	}
	return sum[end + 1] - sum[start + 1 - 1]
}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	alphaToSums = make(map[string][]int)

	fmt.Fscanln(r, &S)

	var q int

	fmt.Fscanln(r, &q)

	for i := 0; i < q; i++ {
		var (
			alpha string
			start, end int
		)
		fmt.Fscanln(r, &alpha, &start, &end)
		fmt.Fprintln(w, Calculate(alpha, start, end))
	}
}