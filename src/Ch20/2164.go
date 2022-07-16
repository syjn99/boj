package main

import (
	"bufio"
	"fmt"
	"os"
)

type Node struct {
	value int
	nextNode *Node
}

type Queue struct {
	size int
	front, back *Node
}

func (queue *Queue) Enqueue(n int) {
	node := Node{n, nil}
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

func (queue *Queue) Dequeue() int {
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

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var N int

	fmt.Fscanln(r, &N)

	queue := Queue{0, nil, nil}
	
	for i := 1; i <= N; i++ {
		queue.Enqueue(i)
	}

	for queue.size > 1 {
		_ = queue.Dequeue()
		num := queue.Dequeue()
		queue.Enqueue(num)
	}

	fmt.Fprintln(w, queue.front.value)
}