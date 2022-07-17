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
	prev, next *Node
}

type Deque struct {
	size int
	front, back *Node
}

func (deque *Deque) PopFront() (val int) {
	val = deque.front.value
	if deque.size == 1 {
		deque.size, deque.front, deque.back = 0, nil, nil
		return
	}
	deque.front = deque.front.next
	deque.front.prev = nil
	deque.size--
	return
}

func (deque *Deque) PopBack() (val int) {
	val = deque.back.value
	if deque.size == 1 {
		deque.size, deque.front, deque.back = 0, nil, nil
		return
	}
	deque.back = deque.back.prev
	deque.back.next = nil
	deque.size--
	return
}

func (deque *Deque) PushBack(val int) {
	node := Node{val, nil, nil}
	if deque.size == 0 {
		deque.size, deque.front, deque.back = 1, &node, &node
		return
	}
	node.prev = deque.back
	deque.back.next = &node
	deque.back = &node
	deque.size++
	return
}

func (deque *Deque) D(isReverse bool) bool {
	if deque.size == 0 {
		return true
	}
	if isReverse {
		_ = deque.PopBack()
	} else {
		_ = deque.PopFront()
	}
	return false
}

func AC(p string, n int, S string) string {
	deque := Deque{0, nil, nil}
	length := len(S)

	S = S[1:(length - 1)]

	if length != 0 {
		numStr := strings.Split(S, ",")
		for _, v := range numStr {
			num, _ := strconv.Atoi(v)
			deque.PushBack(num)
		}
	}

	isReverse := false

	lenP := len(p)

	for i := 0; i < lenP; i++ {
		c := p[i]
		switch c {
		case 'R':
			if isReverse {
				isReverse = false
			} else {
				isReverse = true
			}
		case 'D':
			isErr := deque.D(isReverse)
			if isErr {
				return "error"
			}
		default:
			continue
		}
	
	}
	
	if deque.size == 0 {
		return "[]"
	}

	resultStr := "["

	if isReverse {
		walk := deque.back
		for walk != nil {
			resultStr += strconv.Itoa(walk.value)
			resultStr += ","
			walk = walk.prev
		}

	} else {
		walk := deque.front
		for walk != nil {
			resultStr += strconv.Itoa(walk.value)
			resultStr += ","
			walk = walk.next
		}
	}
	resultStr = strings.TrimRight(resultStr, ",") + "]"
	return resultStr
}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var (
		T, n int
		p, S string
	)

	fmt.Fscanln(r, &T)

	for i := 0; i < T; i++ {
		p, _ = r.ReadString('\n')
		p = strings.TrimSpace(p)
		fmt.Fscanln(r, &n)
		S, _ = r.ReadString('\n')
		S = strings.TrimSpace(S)
		fmt.Fprintln(w, AC(p, n, S))
	}
}