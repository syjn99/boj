package main

import "fmt"

func main() {
	var (
		s string
		cnt int
	)

	fmt.Scanln(&s)

	for i := 0; i < len(s); i++ {
		cnt++
		c := s[i]
		switch {
		case c == 'c':
			if i < len(s) - 1 && (s[i + 1] == '=' || s[i + 1] == '-') {
				i++
				continue
			}
		case c == 'd':
			if i < len(s) - 1 && s[i + 1] == '-' {
				i++
				continue
			} else if i < len(s) - 2 && s[i + 1] == 'z' && s[i + 2] == '=' {
				i = i + 2
				continue
			}
		case c == 'l':
			if i < len(s) - 1 && s[i + 1] == 'j' {
				i++
				continue
			}
		case c == 'n':
			if i < len(s) - 1 && s[i + 1] == 'j' {
				i++
				continue
			}
		case c == 's':
			if i < len(s) - 1 && s[i + 1] == '=' {
				i++
				continue
			}
		case c == 'z':
			if i < len(s) - 1 && s[i + 1] == '=' {
				i++
				continue
			}
		}
	}

	fmt.Println(cnt)
}