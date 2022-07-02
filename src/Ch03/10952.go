package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var a, b int
	r := bufio.NewReader(os.Stdin)
	// w := bufio.NewWriter(os.Stdout)
	for {
		fmt.Fscanf(r, "%d %d\n", &a, &b)
		if a == 0 && b == 0 {
			return
		}
		fmt.Printf("%d\n", a+b)
	}
}