package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var Paper [][]string

func Check(row, col, size int) bool {
	key := Paper[row][col]
	for i := row; i < (row + size); i++ {
		for j := col; j < (col + size); j++ {
			if key != Paper[i][j] {
				return false
			}
		}
	}
	return true
}

func Cutting(row, col, size int) (white, blue int) {
	if !Check(row, col, size) {
		for i := 0; i < 2; i++ {
			for j := 0; j < 2; j++ {
				wNum, bNum := Cutting(row + (i * size/2), col + (j * size/2), size/2)
				white += wNum
				blue += bNum
			}
		}
	} else {
		key := Paper[row][col]
		if key == "1" {
			blue++
		} else {
			white++
		}
	}
	return
} 

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var N int

	fmt.Fscanln(r, &N)

	Paper = make([][]string, N)

	for i := 0; i < N; i++ {
		line, _ := r.ReadString('\n')
		line = strings.TrimSpace(line)
		Paper[i] = strings.Split(line, " ")
	}

	white, blue := Cutting(0, 0, N)

	fmt.Fprintf(w, "%d\n%d\n", white, blue)
}	