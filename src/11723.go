package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var M int

	fmt.Fscanln(r, &M)

	m := make(map[string]bool)

	for i := 0; i < M; i++ {
		S, _ := r.ReadString('\n')
		S = strings.TrimSpace(S)
		opers := strings.Split(S, " ")
		oper := opers[0]
		switch oper {
		case "add":
			m[opers[1]] = true
		case "remove":
			m[opers[1]] = false 
		case "check":
			if m[opers[1]] {
				fmt.Fprintln(w, 1)
			} else {
				fmt.Fprintln(w, 0)
			}
		case "toggle":
			m[opers[1]] = !m[opers[1]]
		case "all":
			for j := 1; j <= 20; j++ {
				m[strconv.Itoa(j)] = true
			}
		case "empty":
			for j := 1; j <= 20; j++ {
				m[strconv.Itoa(j)] = false
			}
		default:
			continue
		}
	}
}