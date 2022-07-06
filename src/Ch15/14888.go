package main

import (
	"bufio"
	"fmt"
	"os"
)

var results []int
var N int

func MaxMin(results []int) (max, min int) {
	max, min = results[0], results[0]
	for _, v := range results {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	return
}

func Compute(index, tempResult int, nums, oper []int) {
	if len(nums) == 0 {
		results = append(results, tempResult)
		return
	}

	temp := tempResult

	for i := 0; i < 4; i++ {
		if oper[i] != 0 {
			oper[i] -= 1
			switch {
			case i == 0:
				temp = tempResult + nums[0]
			case i == 1:
				temp = tempResult - nums[0]
			case i == 2:
				temp = tempResult * nums[0]
			case i == 3:
				temp = tempResult / nums[0]
			}
			Compute(index + 1, temp, nums[1:], oper)
			oper[i] += 1
		}
	}
}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var x int

	fmt.Fscanln(r, &N)
	nums := make([]int, N)
	oper := make([]int, 4)

	for i := 0; i < N; i++ {
		fmt.Fscan(r, &x)
		nums[i] = x
	}

	for i := 0; i < 4; i++ {
		fmt.Fscan(r, &x)
		oper[i] = x
	}

	Compute(0, nums[0], nums[1:], oper)
	max, min := MaxMin(results)
	fmt.Fprintln(w, max)
	fmt.Fprintln(w, min)
}