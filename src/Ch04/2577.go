package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	var a, b, c int
	fmt.Fscanln(r, &a)
	fmt.Fscanln(r, &b)
	fmt.Fscanln(r, &c)

	X := strconv.Itoa(a * b * c)
	arr := make([]int, 10, 10)
	for _, v := range X {
		arr[v - 48]++
	}
	for _, x := range arr {
		fmt.Fprintln(w, x)
	}
	w.Flush()
}