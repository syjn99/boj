package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	fmt.Fscanf(reader, "%d\n", &n)
	for i := 1; i <= n; i++ {
		fmt.Fprintf(writer, "%d\n", i)
	}
	writer.Flush()
}