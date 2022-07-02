package main

import (
	"bufio"
	"fmt"
	"os"
)

type TWODARR [][]rune

func (arrs *TWODARR) ModifySpecifically(N, x, y int) {
	Stars := *arrs
	if N == 3 {
		for i := x; i < x + 3; i++ {
			Stars[i][y] = '*'
			Stars[i][y + 1] = '*'
			Stars[i][y + 2] = '*'
		}
		Stars[x + 1][y + 1] = ' '
	} else {
		unit := N / 3
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				if i == 1 && j == 1 {
					continue
				}
				Stars.ModifySpecifically(unit, x + i * unit, y + j * unit)
				Stars.ModifyBlank(unit, x + i * unit, y + j * unit)
			}
		}
	}
		for i := x; i < N + x; i++ {
		for j := y; j < N + y; j++ {
			if Stars[i][j] == 0 {
				Stars[i][j] = ' '
			}
		}
	}
}

func (arrs *TWODARR) ModifyBlank(N, x, y int) {
	Stars := *arrs
	unit := N / 3
	startX := x + unit
	startY := y + unit
	for i := startX; i < startX + unit; i++ {
		for j := startY; j < startY + unit; j++ {
			Stars[i][j] = ' '
		}
	}
}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	var N int
	fmt.Fscanln(r, &N)
	var Stars TWODARR
	Stars = make([][]rune, N)
	for i := 0; i < len(Stars); i++ {
		Stars[i] = make([]rune, N)
	}

	Stars.ModifySpecifically(N, 0, 0)
	
	for _, v := range Stars {
		for _, c := range v {
			fmt.Fprintf(w, "%c", c)
		}
		fmt.Fprintf(w, "\n")
	}
	w.Flush()
}