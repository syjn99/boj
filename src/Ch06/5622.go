package main

import "fmt"

func Conv(c rune) int {
	switch {
	case c >= 'A' && c <= 'C':
		return 2
	case c >= 'D' && c <= 'F':
		return 3
	case c >= 'G' && c <= 'I':
		return 4
	case c >= 'J' && c <= 'L':
		return 5
	case c >= 'M' && c <= 'O':
		return 6
	case c >= 'P' && c <= 'S':
		return 7
	case c >= 'T' && c <= 'V':
		return 8
	case c >= 'W' && c <= 'Z':
		return 9
	}
	return 0
}

func main() {
	var S string
	fmt.Scanln(&S)
	var time int
	for _, c := range S {
		time += (Conv(c) - 1) + 2
	}

	fmt.Println(time)
}