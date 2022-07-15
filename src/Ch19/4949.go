package main

import (
	"bufio"
	"fmt"
	"os"
)

func isBalanced(S string) bool {
	var stack []rune
	for _, c := range S {
		switch c {
		case '(':
			stack = append(stack, '(')
		case ')':
			if len(stack) == 0 {
				return false
			}
			if stack[len(stack) - 1] != '(' {
				return false
			}
			stack = stack[:len(stack) - 1]
		case '[':
			stack = append(stack, '[')
		case ']':
			if len(stack) == 0 {
				return false
			}
			if stack[len(stack) - 1] != '[' {
				return false
			}
			stack = stack[:len(stack) - 1]
		}
	}
	if len(stack) == 0 {
		return true
	}
	return false
}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var S string
	
	for {
		S, _ = r.ReadString('.')
		if len(S) == 2 {
			break
		}

		if isBalanced(S) {
			fmt.Println("yes")
		} else {
			fmt.Println("no")
		}
	}
}