package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	val int
	prev, next *Node
}

type Deque struct {
	size int
	front, back *Node
}

func (deque *Deque) PushFront(n int) {
	node := Node{n, nil, nil}
	if deque.size == 0 {
		deque.size++
		deque.front, deque.back = &node, &node
		return
	}

	node.next = deque.front
	deque.front.prev = &node
	deque.front = &node
	deque.size++
	return
}

func (deque *Deque) PushBack(n int) {
	if deque.size == 0 {
		deque.PushFront(n)
		return
	}

	node := Node{n, nil, nil}

	node.prev = deque.back
	deque.back.next = &node
	deque.back = &node
	deque.size++
	return
}

func (deque *Deque) PopFront() int {
	if deque.size == 0 {
		return -1
	} else if deque.size == 1 {
		elem := deque.front.val
		deque.size--
		deque.front, deque.back = nil, nil
		return elem
	}
	elem := deque.front.val
	deque.size--
	deque.front = deque.front.next
	deque.front.prev = nil
	return elem
}

func (deque *Deque) PopBack() int {
	if deque.size <= 1 {
		return deque.PopFront()
	}
	elem := deque.back.val
	deque.size--
	deque.back = deque.back.prev
	deque.back.next = nil
	return elem
}

func (deque Deque) Size() int {
	return deque.size
}

func (deque Deque) Empty() int {
	if deque.size == 0 {
		return 1
	}
	return 0
}

func (deque Deque) Front() int {
	if deque.size == 0 {
		return -1
	}
	return deque.front.val
}

func (deque Deque) Back() int {
	if deque.size == 0 {
		return -1
	}
	return deque.back.val
}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var N int
	fmt.Fscanln(r, &N)

	deque := Deque{0, nil, nil}

	for i := 0; i < N; i++ {
		S, _ := r.ReadString('\n')
		opers := strings.Split(S, " ")
		if len(opers) == 2 {
			opers[1] = opers[1][:len(opers[1]) - 1]
		} else {
			opers[0] = opers[0][:len(opers[0]) - 1]
		}
		switch opers[0] {
		case "push_front":
			num, _ := strconv.Atoi(opers[1])
			deque.PushFront(num)
		case "push_back":
			num, _ := strconv.Atoi(opers[1])
			deque.PushBack(num)
		case "pop_front":
			fmt.Fprintln(w, deque.PopFront())
		case "pop_back":
			fmt.Fprintln(w, deque.PopBack())
		case "size":
			fmt.Fprintln(w, deque.Size())
		case "empty":
			fmt.Fprintln(w, deque.Empty())
		case "front":
			fmt.Fprintln(w, deque.Front())
		case "back":
			fmt.Fprintln(w, deque.Back())
		default:
			continue
		}
	}	
}