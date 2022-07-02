package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	str, _ := r.ReadString('\n')
	words := strings.Split(str, " ")
	wordCnt := 0

	for i := range words {
		if words[i] != "\n" && words[i] != "" {
			wordCnt++
		}
	}
	
	fmt.Fprintln(w, wordCnt)

	w.Flush()
}