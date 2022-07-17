package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	r *bufio.Reader
	w *bufio.Writer
)

type Deque struct {
	item []string
}

func (deque *Deque) PopFront() {
	deque.item = deque.item[1:]
}

func (deque *Deque) PopBack() {
	deque.item = deque.item[:(len(deque.item) - 1)]
}

func AC(p string, S string) {
	deque := Deque{}
	if len(S) == 0 {
		deque.item = []string{}
	} else {
		deque.item = strings.Split(S, ",")
	}

	isReversed := false

	for _, c := range p {
		switch c {
		case 'R':
			isReversed = !isReversed
		case 'D':
			if len(deque.item) == 0 {
				fmt.Fprintln(w, "error")
				return
			}
			if isReversed {
				deque.PopBack()
			} else {
				deque.PopFront()
			}
		default:
			continue
		}
	}

	if len(deque.item) == 0 {
		fmt.Fprintln(w, "[]")
		return
	}

	fmt.Fprint(w, "[")

	if isReversed {
		for i := (len(deque.item) - 1); i >= 0; i-- {
			fmt.Fprintf(w, "%s", deque.item[i])
			if i != 0{
				fmt.Fprintf(w, ",")
			}
		}
	} else {
		fmt.Fprint(w, strings.Join(deque.item, ","))
	}

	fmt.Fprint(w, "]\n")
}

func main() {
	r = bufio.NewReader(os.Stdin)
	w = bufio.NewWriter(os.Stdout)
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
		AC(p, S[1:(len(S) - 1)])
	}
}