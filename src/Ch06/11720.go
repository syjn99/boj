package main

import "fmt"

func main() {
	var (
		N, result int
		nums string
	)
	fmt.Scanln(&N)
	fmt.Scanln(&nums)

	result = 0
	for i := 0; i < N; i++ {
		result += (int(nums[i]) - 48)
	}

	fmt.Println(result)
}