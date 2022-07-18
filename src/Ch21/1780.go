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

func Cutting(row, col, size int) (minus, zero, plus int) {
	if !Check(row, col, size) {
		for i := 0; i < 3 ;i++ {
			for j := 0; j < 3; j++ {
				mNum, zNum, pNum := Cutting(row + (i * size/3), col + (j * size/3), size/3)
				minus += mNum
				zero += zNum
				plus += pNum
			}
		}
	} else {
		key := Paper[row][col]
		switch key {
		case "-1":
			minus++
		case "0":
			zero++
		case "1":
			plus++
		default:
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

	minus, zero, plus := Cutting(0, 0, N)

	fmt.Fprintf(w, "%d\n%d\n%d\n", minus, zero, plus)
}	