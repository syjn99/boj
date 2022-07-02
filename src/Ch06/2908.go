package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func flip(s string) string {
	arr := make([]byte, 3)
	arr[0] = s[2]
	arr[1] = s[1]
	arr[2] = s[0]
	return string(arr)
}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	var a, b string
	fmt.Fscan(r, &a)
	fmt.Fscan(r, &b)
	
	A, _ := strconv.Atoi(flip(a))
	B, _ := strconv.Atoi(flip(b))

	if A > B {
		fmt.Fprintln(w, A)
	} else {
		fmt.Fprintln(w, B)
	}


	w.Flush()
}