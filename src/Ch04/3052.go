package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	arr := make([]int, 10, 10)
	m := make(map[int]bool)

	var x int
	count := 0

	for i := 0; i < 10; i++ {
		fmt.Fscanln(r, &x)
		arr[i] = x % 42
		_, ok := m[arr[i]]
		if ok == false {
			m[arr[i]] = true
			count++
		}
	}

	fmt.Fprintln(w, count)

	w.Flush()
}