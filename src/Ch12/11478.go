package main

import "fmt"

func main() {
	var (
		S string
		cnt int
	)

	fmt.Scanln(&S)
	m := make(map[string]bool)
	for i := 0; i < len(S); i++ {
		for j := i + 1; j <= len(S); j++ {
			if !m[S[i:j]] {
				m[S[i:j]] = true
				cnt++
			}
		}
	}
	fmt.Println(cnt)
}