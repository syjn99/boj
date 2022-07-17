package main

import (
	"bufio"
	"fmt"
	"os"
)

type Node struct {
	index, priority int
	next *Node
}

type Queue struct {
	size int
	front, back *Node
}

func (queue *Queue) Enqueue(index, priority int) {
	node := Node{index, priority, nil}
	if queue.back == nil {
		queue.front, queue.back = &node, &node
		queue.size++
		return
	}

	queue.back.next = &node
	queue.back = &node
	queue.size++
	return
}

func (queue *Queue) Dequeue() (index, priority int) {
	if queue.size == 1 {
		node := queue.front
		index, priority = node.index, node.priority
		queue.size--
		queue.front, queue.back = nil, nil
		return
	}

	index, priority = queue.front.index, queue.front.priority
	queue.front = queue.front.next
	queue.size--
	return
}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var T, N, M, x int

	
	fmt.Fscanln(r, &T)
	
	for i := 0; i < T; i++ {
		queue := Queue{0, nil, nil}
		fmt.Fscanf(r, "%d %d\n", &N, &M)
		for j := 0; j < N; j++ {
			fmt.Fscanf(r, "%d ", &x)
			queue.Enqueue(j, x)
		}
		fmt.Fscanf(r, "")

		cnt := 0

		for {
			walk, priority, DeqEnq := queue.front, queue.front.priority, false
			for walk.next != nil {
				walk = walk.next
				if priority < walk.priority {
					DeqEnq = true
					ind, pri := queue.Dequeue()
					queue.Enqueue(ind, pri)
					break
				}
			}
			if !DeqEnq {
				cnt++
				ind, _ := queue.Dequeue()
				if ind == M {
					fmt.Fprintln(w, cnt)
					break
				}
			}
		}
		
	}
}