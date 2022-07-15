package main

import (
	"bufio"
	"fmt"
	"os"
)

type Stack struct {
	stackArr []int
}

func (stack *Stack) Push(n int) {
	stack.stackArr = append(stack.stackArr, n)
}

func (stack *Stack) Pop() {
	stack.stackArr = stack.stackArr[:len(stack.stackArr) - 1]
}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var K, x int
	stack := Stack{}

	fmt.Fscanln(r, &K)

	for i := 0; i < K; i++ {
		fmt.Fscanln(r, &x)
		if x == 0 {
			stack.Pop()
		} else {
			stack.Push(x)
		}
	}
	arr := stack.stackArr
	sum := 0
	for _, v := range arr {
		sum += v
	}
	fmt.Fprintln(w, sum)
}