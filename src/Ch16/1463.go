package main

import "fmt"

var (
	memo []int
	memoBool []bool
)

func MinArr(arr []int) int {
	min := arr[0]

	for _, v := range arr {
		if v < min {
			min = v
		}
	}

	return min
}

func MakeOne(X int) int {
	if X == 1 {
		memo[1] = 0
		memoBool[1] = true
	} else if X <= 3 {
		memo[X] = 1
		memoBool[X] = true
	}

	if memoBool[X] {
		return memo[X]
	}

	var tempArr []int

	if X % 3 == 0 {
		tempArr = append(tempArr, MakeOne(X / 3))
	}
	if X % 2 == 0 {
		tempArr = append(tempArr, MakeOne(X / 2))
	}
	tempArr = append(tempArr, MakeOne(X - 1))
	memo[X] = MinArr(tempArr) + 1
	memoBool[X] = true
	return MakeOne(X)
}

func main() {
	var X int

	fmt.Scan(&X)
	memo = make([]int, X + 1)
	memoBool = make([]bool, X + 1)
	fmt.Println(MakeOne(X))
}