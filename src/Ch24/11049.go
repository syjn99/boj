package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type RC struct {
	row, col int
}

func Update(com *[][]int, mat *[][]RC, ind, row, col int) {
	COM, MAT := *com, *mat

	// init setting
	COM[ind][ind] = 0
	MAT[ind][ind] = RC{row, col}


	for i := ind-1; i >= 0; i-- {
		if i == ind-1 {
			COM[i][ind] = MAT[i][i].row * MAT[i][i].col * MAT[ind][ind].col
			MAT[i][ind] = RC{MAT[i][i].row, MAT[ind][ind].col}
			continue
		}
		min := math.MaxInt
		// i부터 j까지, j+1부터 ind까지
		for j := i; j < ind; j++ {
			comp := COM[i][j] + COM[j+1][ind]
			comp += MAT[i][j].row * MAT[i][j].col * MAT[j+1][ind].col
			if comp < min {
				min = comp
				MAT[i][ind] = RC{MAT[i][j].row, MAT[j+1][ind].col}
			}
		}
		COM[i][ind] = min
	}
}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var N, row, col int

	fmt.Fscanln(r, &N)

	com, mat := make([][]int, N), make([][]RC, N)

	for i := 0; i < N; i++ {
		com[i] = make([]int, N)
		mat[i] = make([]RC, N)
	}

	for i := 0; i < N; i++ {
		fmt.Fscanln(r, &row, &col)
		Update(&com, &mat, i, row, col)
	}
	fmt.Fprintln(w, com[0][N-1])
}