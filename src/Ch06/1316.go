package main

import "fmt"

func isGroup(s string) int {
	m := make(map[rune]int)
	for i, c := range s {
		if m[c] == 0 {
			m[c] = 1
			if i >= 1 {
				m[rune(s[i - 1])] = 2
			}
		} else if m[c] == 2 {
			return 0
		}
	}
	return 1
}

func main() {
	var (
		N, cnt int
		s string
	)

	fmt.Scanln(&N)

	for i := 0; i < N; i++ {
		fmt.Scanln(&s)
		cnt += isGroup(s)
	}

	fmt.Println(cnt)
}