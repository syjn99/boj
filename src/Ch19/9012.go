package main

import (
	"bufio"
	"fmt"
	"os"
)

func isVPS(S string) bool {
	stackNum := 0
	for _, c := range S {
		if c == '(' {
			stackNum++
		} else if c == ')' {
			stackNum--
			if stackNum < 0 {
				return false
			}
		}
	}
	if stackNum == 0 {
		return true
	}
	return false
}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var (
		T int
		S string
	)

	fmt.Fscanln(r, &T)

	for i := 0; i < T; i++ {
		fmt.Fscanln(r, &S)
		if isVPS(S) {
			fmt.Fprintln(w, "YES")
		} else {
			fmt.Fprintln(w, "NO")
		}
	}
}