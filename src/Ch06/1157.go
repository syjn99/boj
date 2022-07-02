package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	var (
		S string
		iterated int
	)

	fmt.Fscanln(r, &S)

	m := make(map[int]int)

	for _, c := range S {
		if c >= 97 {
			c -= 32
		}
		m[int(c)]++
	}

	maxChar := 65
	maxCnt := m[65]
	iterated = 0
	

	for i := 66; i <= 90; i++ {
		if m[i] > maxCnt {
			iterated = 0
			maxChar = i
			maxCnt = m[i]
		} else if m[i] == maxCnt {
			iterated = 1
		}
	}

	if iterated == 1 {
		fmt.Fprintln(w, "?")
	} else {
		fmt.Fprintf(w, "%c\n", maxChar)
	}

	w.Flush()
}