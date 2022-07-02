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

	var (
		N, M int
		name string
		input string
	)

	fmt.Fscanf(r, "%d %d\n", &N, &M)

	nameToNum := make(map[string]int)
	numToName := make(map[int]string)

	for i := 0; i < N; i++ {
		fmt.Fscanln(r, &name)
		nameToNum[name] = i + 1
		numToName[i + 1] = name
	}

	for i := 0; i < M; i++ {
		fmt.Fscanln(r, &input)
		if input[0] <= 57 {
			num, _ := strconv.Atoi(input)
			fmt.Fprintln(w, numToName[num])
		} else {
			fmt.Fprintln(w, nameToNum[input])
		}
	}

	w.Flush()
}