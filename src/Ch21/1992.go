package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var Video [][]string

func Check(row, col, size int) bool {
	key := Video[row][col]
	for i := row; i < (row + size); i++ {
		for j := col; j < (col + size); j++ {
			if key != Video[i][j] {
				return false
			}
		}
	}
	return true
}

func QuadTree(row, col, size int) (result string) {
	if !Check(row, col, size) {
		result = "("
		for i := 0; i < 2; i++ {
			for j := 0; j < 2; j++ {
				result += QuadTree(row + (i * size/2), col + (j * size/2), size/2)
			}
		}
		result += ")"
	} else {
		result = Video[row][col]
	}
	return
} 

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var N int

	fmt.Fscanln(r, &N)

	Video = make([][]string, N)

	for i := 0; i < N; i++ {
		line, _ := r.ReadString('\n')
		line = strings.TrimSpace(line)
		Video[i] = strings.Split(line, "")
	}

	fmt.Fprintln(w, QuadTree(0, 0, N))
}	