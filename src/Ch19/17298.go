package main

import (
	"bufio"
	"fmt"
	"os"
)

type Elem struct {
	value, index int
}

type Stack struct {
	stackArr []Elem
	size int
}

func newStack() *Stack {
	stack := Stack{}
	stack.size = - 1
	stack.stackArr = make([]Elem, 0)
	return &stack
}

func (stack *Stack) Push(elem Elem) {
	stack.stackArr = append(stack.stackArr, elem)
	stack.size++
}

func (stack *Stack) Pop() Elem {
	if stack.size == -1 {
		return Elem{-1, -1}
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
	return stack.stackArr[stack.size].value
}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var N, x int

	fmt.Fscanln(r, &N)
	result := make([]int, N)
	result[N - 1] = -1
	stack := newStack()
	for i := 0; i < N; i++ {
		fmt.Fscan(r, &x)
		
		top := stack.Top()
		if top == -1 {
			stack.Push(Elem{x, i})
			continue
		}
		for top < x {
			elem := stack.Pop()
			result[elem.index] = x
			top = stack.Top()
			if top == -1 {
				break
			}
		}
		stack.Push(Elem{x, i})
	}

	for stack.size > -1 {
		elem := stack.Pop()
		result[elem.index] = -1
	}

	for _, v := range result {
		fmt.Fprintf(w, "%d ", v)
	}
	fmt.Fprintf(w, "\n")
}