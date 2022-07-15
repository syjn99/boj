package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

func (stack Stack) Size() int {
	return stack.size + 1
}

func (stack Stack) Empty() int {
	if stack.size == -1 {
		return 1
	}
	return 0
}

func (stack Stack) Top() int {
	if stack.size == -1 {
		return -1
	}
	return stack.stackArr[stack.size]
}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	stack := newStack()

	var N int

	fmt.Fscanln(r, &N)
	for i := 0; i < N; i++ {
		oper, _ := r.ReadString('\n')
		slice := strings.Split(oper[:(len(oper) - 1)], " ")
		switch slice[0] {
		case "push":
			num, _ := strconv.Atoi(slice[1])
			stack.Push(num)
		case "pop":
			elem := stack.Pop()
			fmt.Fprintln(w, elem)
		case "size":
			size := stack.Size()
			fmt.Fprintln(w, size)
		case "empty":
			isEmpty := stack.Empty()
			fmt.Fprintln(w, isEmpty)
		case "top":
			top := stack.Top()
			fmt.Fprintln(w, top)
		default:
			continue
		}
	}
}