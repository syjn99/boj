// 나머지가 같다는 것은, 원소들을 서로 뺐을 때 그 뺀 값들이 공통인 약수를 가진다는 것! 그 중 최대 공약수를 찾아서 최대 공약수의 약수들을 출력했다.

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func PrintCDs(x int, w *bufio.Writer) {
	for i := 2; i <= x / 2; i++ {
		if x % i == 0 {
			fmt.Fprintf(w, "%d ", i)
		}
	}
	fmt.Fprintf(w, "%d\n", x)
}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	var N, x int
	fmt.Fscanln(r, &N)
	arr := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscanln(r, &x)
		arr[i] = x
	}
	sort.Ints(arr)

	for {
		subs := make([]int, len(arr) - 1)
		var vals []int
		m := make(map[int]bool)
		for i := 1; i < len(arr); i++ {
			sub := arr[i] - arr[i - 1]
			if sub <= 3 {
				fmt.Fprintln(w, sub)
				w.Flush()
				return
			}
			subs[i - 1] = sub
		}
		sort.Ints(subs)
		for i := 0; i < len(subs); i++ {
			val := subs[i]
			if m[val] == false {
				vals = append(vals, val)
				for j := i + 1; j < len(subs); j++ {
					if subs[j] % val == 0 {
						m[subs[j]] = true
					}
				}
			}
		}
		if len(vals) == 1 {
			PrintCDs(subs[0], w)
			w.Flush()
			return
		}
		arr = vals
	}
}