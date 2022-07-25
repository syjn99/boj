package main

import (
	"bufio"
	"fmt"
	"os"
)

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var N, K int

	fmt.Fscan(r, &N, &K)

	left, right := 1, N*N

	for {
		mid := (left + right) / 2
		leftSide, rightSide := 0, 0
		for i := 1; i <= N; i++ {
			temp := mid / i
			leftSide += Min(temp, N)
			rightSide += Min(temp, N)
			if Min(temp, N) == temp && mid % i == 0 {
				leftSide--
			}
		}
		if leftSide == rightSide {
			if leftSide >= K {
				right = mid-1
			} else {
				left = mid+1
			}
			continue
		}
		leftSide++
		if leftSide <= K && rightSide >= K {
			fmt.Fprintln(w, mid)
			break
		} else if leftSide > K {
			right = mid-1
		} else if rightSide < K {
			left = mid+1
		}
	}
}

	/*  1 2 3 4 5 
			2 4 6 8 10
			3 6 9 12 15
			4 8 12 16 20
			5 10 15 20 25*/
			// N = 5, K = 16 답 10
			// try1: left = 1, right = 25, mid = 13 => 13/1 비교-> 13개까지는 가능, 근데 5개밖에 없 -> 5개, 13/2 => 6개까지는 가능, 근데 5개밖에 없. -> 5개, total 10개.
			// 13/3 => 4개, total 14개 13/4 => 3개, 13/5 => 2개 => 19개. 19 < x <= 19
			// try1 결과 K = 16 < 19이므로, 숫자가 더 작아져야 함. 
			// try2: left = 1, right = 12, mid = 6. 5 + 2 + 1 + 1 + 1 = 10, 5 + 3 + 2 + 1 + 1 = 12개. 숫자가 더 커져야 함. 10 < x <= 12
			// try3: left = 7, right = 12, mid = 8. 5 + 3 + 2 + 1 + 1 = 12, 5 + 4 + 2 + 2 + 1 = 14개. 숫자가 더 커져야 함.
			// try4: left = 9, right = 12, mid = 10. 5 + 4 + 3 + 2 + 1 = 15(10보다 작은 숫자들) 5 + 5 + 3 + 2 + 2 = 17개. (10보다 작거나 같은 숫자들.) 숫자가 더 작아져야 함.
			// try5: left = 9, right = 11, mid = 10. 숫자가 더 작아져야 함.
			// try6: left = 9, 

			// N = 3, K = 7, 답 6
			// try1: left = 1, right = 9, mid = 5. 3 + 2 + 1 = 6. 6<= <=6 (5는 행렬에 없는 숫자. 따라서 수를 키워야 함.)
			// try2: left = 6, right = 9, mid = 7. 8. 숫자가 너무 크다
			// try3: left = 6, right = 6. mid = 6. 6 < x <= 8 ok


		// 1 x 1 => 1
		// 2 x 2 => 1 2 2 4
		// 3 x 3 => 1 2 2 3 3 4 6 6 9
		// 4 x 4 => 1 2 2 3 3 4 4 4 6 6 8 8 9 12 12 16
		// 5 x 5 => 1 2 2 3 3 4 4 4 5 5 6 6 8 8 9 10 10 12 12 15 15 16 20 20 25