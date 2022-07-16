package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	value int
	nextNode *Node
}

type Queue struct {
	size int
	front, back *Node
}

var queue Queue

func newQueue() Queue {
	queue := Queue{}
	queue.size = 0
	queue.front, queue.back = nil, nil
	return queue 
}

func (queue *Queue) Push(n int) {
	node := Node{}
	node.value = n
	node.nextNode = nil
	if queue.size == 0 {
		queue.front, queue.back = &node, &node
		queue.size++
		return
	}
	back := queue.back
	back.nextNode = &node
	queue.back = &node
	queue.size++
	return
}

func (queue *Queue) Pop() int {
	if queue.size == 0 {
		return -1
	} else if queue.size == 1 {
		val := queue.front.value
		queue.front, queue.back = nil, nil
		queue.size--
		return val
	}

	front := queue.front
	val := front.value
	queue.front = front.nextNode
	queue.size--
	return val
}

func (queue Queue) Size() int {
	return queue.size
}

func (queue Queue) Empty() int {
	if queue.size == 0 {
		return 1
	} else {
		return 0
	}
}

func (queue Queue) Front() int {
	if queue.size == 0 {
		return -1
	}
	return queue.front.value
}

func (queue Queue) Back() int {
	if queue.size == 0 {
		return -1
	}
	return queue.back.value
}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var (
		N int
	)

	fmt.Fscanln(r, &N)

	queue = newQueue()

	for i := 0; i < N; i++ {
		S, _ := r.ReadString('\n')
		opers := strings.Split(S, " ")
		oper := opers[0]
		if len(opers) == 1 {
			oper = oper[:len(oper) - 1]
		}
		switch oper {
		case "push":
			num, _ := strconv.Atoi(opers[1][:len(opers[1]) - 1])
			queue.Push(num)
		case "pop":
			fmt.Fprintln(w, queue.Pop())
		case "size":
			fmt.Fprintln(w, queue.Size())
		case "empty":
			fmt.Fprintln(w, queue.Empty())
		case "front":
			fmt.Fprintln(w, queue.Front())
		case "back":
			fmt.Fprintln(w, queue.Back())
		default:
			continue
		}
	}
}