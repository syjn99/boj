package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	var N, x, cnt int
	
	fmt.Fscanln(r, &N)

	arr := make([]int, N)
	arrSorted := make([]int, N)

	isWritten := make(map[int]bool)
	mapFunc := make(map[int]int)

	for i := 0; i < N; i++ {
		fmt.Fscan(r, &x)
		arr[i] = x
		arrSorted[i] = x
		if isWritten[x] == false {
			cnt++
			isWritten[x] = true
		}
	}

	sort.Ints(arrSorted)

	squashNum := 0
	mapFunc[arrSorted[0]] = squashNum

	for i := 1; i < N; i++ {
		if arrSorted[i] == arrSorted[i - 1] {
			continue
		}
		squashNum++
		mapFunc[arrSorted[i]] = squashNum
	}

	for i := 0; i < N; i++ {
		fmt.Fprintf(w, "%d ", mapFunc[arr[i]])		
	}
	
	fmt.Fprintf(w, "\n")

	w.Flush()
}