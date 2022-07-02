package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var a, b int
	r := bufio.NewReader(os.Stdin)
	for {
		n, _ := fmt.Fscanf(r, "%d %d\n", &a, &b)
		if n == 0 {
			return
		}
		fmt.Printf("%d\n", a+b)
	}
}