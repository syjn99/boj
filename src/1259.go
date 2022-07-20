package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func isPalindrome(N string, start, end int) bool {
	if start >= end {
		return true
	}
	if N[start] != N[end] {
		return false
	} else {
		if isPalindrome(N, start+1, end-1) {
			return true
		} else {
			return false
		}
	}
}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var N string

	for {
		fmt.Fscanln(r, &N)
		if num, _ := strconv.Atoi(N); num == 0 {
			break
		}
		if isPalindrome(N, 0, len(N) - 1) {
			fmt.Fprintln(w, "yes")
		} else {
			fmt.Fprintln(w, "no")
		}
	}
}