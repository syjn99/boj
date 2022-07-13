package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)


type Meeting struct {
	start, end int
}

var Meetings []Meeting

func Filter(Meetings []Meeting, f func(Meeting) bool) []Meeting {
	returnSlice := make([]Meeting, 0)
	for _, v := range Meetings {
		if f(v) {
			returnSlice = append(returnSlice, v)
		}
	}

	return returnSlice
}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var N, start, end int

	fmt.Fscanln(r, &N)
	
	Meetings = make([]Meeting, 0)

	for i := 0; i < N; i++ {
		fmt.Fscanln(r, &start, &end)
		Meetings = append(Meetings, Meeting{start, end})
	}

	sort.Slice(Meetings, func(i, j int) bool {
		first, second := Meetings[i], Meetings[j]
		if first.end == second.end {
			return first.start < second.start
		}
		return first.end < second.end
	})

	cnt, firstEnd := 1, Meetings[0].end

	for i := 1; i < N; i++ {
		meet := Meetings[i]
		if meet.start >= firstEnd {
			cnt++
			firstEnd = meet.end
		}
	}
	fmt.Fprintln(w, cnt)
}