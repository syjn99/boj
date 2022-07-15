package main

import (
	"bufio"
	"fmt"
	"os"
)

type Stack struct {
	stackArr []int
	size int
}

func newStack() *Stack {
	stack := Stack{}
	stack.stackArr = make([]int, 0)
	stack.size = -1
	return &stack
}

func (stack *Stack) Push(n int) {
	stack.stackArr = append(stack.stackArr, n)
	stack.size++
}

func (stack *Stack) Pop() int {
	if stack.size == -1 {
		return -1
	}
	elem := stack.stackArr[stack.size]
	stack.stackArr = stack.stackArr[:stack.size]
	stack.size--
	return elem
}

func (stack Stack) Top() int {
	if stack.size == -1 {
		return -1
	}
	return stack.stackArr[stack.size]
}

func isPossible(arr []int) []string {
	var (
		result []string
		cnt int
	)
	cnt = 1
	stack := newStack()

	for _, v := range arr {
		for stack.Top() < v {
			stack.Push(cnt)
			cnt++
			result = append(result, "+")
		}
		if stack.Top() != v {
			return []string{"NO"}
		}
		stack.Pop()
		result = append(result, "-")
	}
	return result
}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var n, x int

	fmt.Fscanln(r, &n)
	arr := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Fscanln(r, &x)
		arr[i] = x
	}

	resultString := isPossible(arr)

	for _, v := range resultString {
		fmt.Fprintln(w, v)
	}
}