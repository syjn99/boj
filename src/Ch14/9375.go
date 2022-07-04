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
	var (
		T, n int
	)
	fmt.Fscanln(r, &T)
	for i := 0; i < T; i++ {
		closet := make(map[string]int)
		var types []string
		fmt.Fscanln(r, &n)
		for j := 0; j < n; j++ {
			strs, _ := r.ReadString('\n')
			strs = strs[:len(strs) - 1]
			strArr := strings.Split(strs, " ")
			types = append(types, strArr[1])
			closet[strArr[1]]++
		}
		var closetArr []int
		for _, v := range closet {
			closetArr = append(closetArr, v)
		}
		comb := 1
		for _, v := range closetArr {
			comb *= (v + 1)
		}
		comb--
		fmt.Fprintln(w, comb)
	}
	w.Flush()
}