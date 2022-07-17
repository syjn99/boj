package main

import (
	"bufio"
	"fmt"
	"os"
)

type Node struct {
	value, index int
	prev, next *Node
}

type Deque struct {
	size int
	front, back *Node
}

func (deque *Deque) PopFront() (value, index int) {
	value, index = deque.front.value, deque.front.index
	if deque.size == 1 {
		deque.size--
		deque.front, deque.back = nil, nil
		return
	}
	deque.front = deque.front.next
	deque.front.prev = nil
	deque.size--
	return
}

func (deque *Deque) PopBack() (value, index int) {
	value, index = deque.back.value, deque.back.index
	if deque.size == 1 {
		deque.size--
		deque.front, deque.back = nil, nil
		return
	}
	deque.back = deque.back.prev
	deque.back.next = nil
	deque.size--
	return
}

func (deque *Deque) PushFront(value, index int) {
	node := Node{value, index, nil, nil}
	if deque.size == 0 {
		deque.front, deque.back = &node, &node
		deque.size++
		return
	}
	
	node.next = deque.front
	deque.front.prev = &node
	deque.front = &node
	deque.size++
	return
}

func (deque *Deque) PushBack(value, index int) {
	node := Node{value, index, nil, nil}
	if deque.size == 0 {
		deque.front, deque.back = &node, &node
		deque.size++
		return
	}
	
	node.prev = deque.back
	deque.back.next = &node
	deque.back = &node
	deque.size++
	return
}

func (deque *Deque) Left() {
	value, _ := deque.PopFront()
	deque.PushBack(value, deque.back.index + 1)
}

func (deque *Deque) Right() {
	value, _ := deque.PopBack()
	deque.PushFront(value, deque.front.index - 1)
}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var N, M, x int

	fmt.Fscanln(r, &N, &M)

	deque := Deque{0, nil, nil}

	for i := 1; i <= N; i++ {
		deque.PushBack(i, i)
	}

	cnt := 0

	for i := 0; i < M; i++ {
		fmt.Fscan(r, &x)
		walk, index := deque.front, 0
		for walk != nil {
			if walk.value == x {
				index = walk.index
				break
			}
			walk = walk.next
		}
		frontIndex, backIndex := deque.front.index, deque.back.index
		mid := (frontIndex + backIndex) / 2
		// cnt 세주기
		if index <= mid {
			for j := 0; j < (index - frontIndex); j++ {
				cnt++
				deque.Left()
			}
			// 왼쪽으로 한 칸 이동하는 거 수행. 얼만큼? index - frontIndex만큼
		} else {
			for j := 0; j < (backIndex - index + 1); j++ {
				cnt++
				deque.Right()
			}
			// 오른쪽으로 한 칸 이동하는 거 수행. 얼만큼? backIndex - index + 1만큼
		}
		_, _ = deque.PopFront()
		// 이렇게 되면 맨 처음에 우리가 원하는 친구가 있을 것이므로, PopFront 해준다.
	}
	// 4 5 7 8 9 10 2
	fmt.Fprintln(w, cnt)
}