package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Member struct {
	age int
	name string
	sequence int
}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	var (
		N, age int
		name string
	)
	fmt.Fscanln(r, &N)
	arr := make([]Member, N)
	for i := 0; i < N; i++ {	
		fmt.Fscan(r, &age)
		fmt.Fscan(r, &name)
		member := Member{age, name, i}
		arr[i] = member
	}
	sort.Slice(arr, func(i, j int) bool {
		if arr[i].age != arr[j].age {
			return arr[i].age < arr[j].age
		} else {
			return arr[i].sequence < arr[j].sequence
		}
	})
	for _, member := range arr {
		fmt.Fprintf(w, "%d %s\n", member.age, member.name)
	}
	w.Flush()
}