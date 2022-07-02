package main

import "fmt"

func main() {
	var word string
	fmt.Scanln(&word)

	for i := 97; i <= 122; i++ {
		pos := -1
		for index, c := range word {
			if rune(i) == c {
				pos = index
				break
			}
		}
		fmt.Printf("%d ", pos)
	}
	fmt.Printf("\n")
}