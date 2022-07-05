package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var (
	N, M int
	seqs []string
)

func Sequence(seq string, M int, used []int, w *bufio.Writer) {
	usedMap := make(map[int]bool)
	for _, v := range used {
		usedMap[v] = true
	}
	var unused []int
	for i := 1; i <= N; i++ {
		if !usedMap[i] {
			unused = append(unused, i)
		}
	}
	tempSeqs := make(map[string][]int)
	for _, v := range unused {
		temp := seq
		temp = fmt.Sprintf("%s %d", temp, v)
		seqs = append(seqs, temp)
		for _, v1 := range used {
			tempSeqs[temp] = append(tempSeqs[temp], v1)
		}
		tempSeqs[temp] = append(tempSeqs[temp], v)
	}
	if M == 1 {
		return
	} else {
		for k, _ := range tempSeqs {
			Sequence(k, M - 1, tempSeqs[k], w)
		}
	}

}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	fmt.Fscanln(r, &N, &M)
	Sequence("", M, []int{}, w)
	sort.Strings(seqs)
	for _, s := range seqs {
		if len(s) == (2 * M) {
			fmt.Fprintln(w, s[1:])
		}
	}
	w.Flush()
}