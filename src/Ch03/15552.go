package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var T, a, b int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	fmt.Fscanf(reader, "%d\n", &T)
	
	for i := 0; i < T; i++ {
		fmt.Fscanf(reader, "%d %d\n", &a, &b)
		fmt.Fprintf(writer, "%d\n", a+b)
	}
	writer.Flush()
}