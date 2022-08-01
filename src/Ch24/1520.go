package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	currRow, currCol, M, N int
	Map, dp [][]int
)

func Update(row, col, delta int) {
	// 만약 업데이트하고자 하는 값이 아직 채워지지 않은 값이라면, 아무 것도 하지 않는다.
	if row == currRow && col > currCol {
		return
	}

	dp[row][col] += delta

	// 위부터 비교
	if row != 0 && Map[row][col] > Map[row-1][col] {
		Update(row-1, col, delta)
	}
	if col != 0 && Map[row][col] > Map[row][col-1] {
		Update(row, col-1, delta)
	}
	if row != currRow {
		if row+1 != currRow || col <= currCol {
			if Map[row][col] > Map[row+1][col] {
				Update(row+1, col, delta)
			}
		}
	}
	if col != N-1 {
		if row != currRow || col <= currCol {
			if Map[row][col] > Map[row][col+1] {
				Update(row, col+1, delta)
			}
		}
	}
}

func Fill(row, col int) {
	if row == 0 && col == 0 {
		dp[row][col] = 1
		return
	}

	if row == 0 {
		if Map[row][col-1] > Map[row][col] {
			dp[row][col] = dp[row][col-1]
		} else {
			dp[row][col] = 0
		}
		return
	} else if col == 0 {
		if Map[row-1][col] > Map[row][col] {
			dp[row][col] = dp[row-1][col]
		} else {
			dp[row][col] = 0
		}
		return
	}

	left, top, center := Map[row][col-1], Map[row-1][col], Map[row][col]

	// 위, 왼 모두 해당 값보다 클 때 => 두 값을 더 해준다.
	if left > center && top > center {
		dp[row][col] = dp[row][col-1] + dp[row-1][col]
		return
	} else if top > center { // 위만 클 때, 위의 dp 값을 더해준다. 왼이 더 작으면 그 값부터 업데이트해준다.
		dp[row][col] = dp[row-1][col]
		if left < center {
			Update(row, col-1, dp[row][col])
		}
		return
	} else if left > center { // 왼만 클 때, 왼의 dp 값을 더해준다. 위가 더 작으면 그 값부터 업데이트해준다.
		dp[row][col] = dp[row][col-1]
		if top < center {
			Update(row-1, col, dp[row][col])
		}
		return
	} // 위와 왼 모두 센터보다 크지 않을 때, 갈 수 있는 방법은 없으므로 0이다.
	return
}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var x int

	fmt.Fscanln(r, &M, &N)

	Map = make([][]int, M)
	dp = make([][]int, M)

	for i := 0; i < M; i++ {
		currRow = i
		Map[i] = make([]int, N)
		dp[i] = make([]int, N)
		for j := 0; j < N; j++ {
			// fmt.Printf("current pos: %d, %d\n", i, j)
			currCol = j
			fmt.Fscan(r, &x)
			Map[i][j] = x
			Fill(i, j)
		}
	}
	fmt.Fprintln(w, dp[M-1][N-1])
}