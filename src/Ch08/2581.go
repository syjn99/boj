package main

import "fmt"

func isPrime(x int) int {
	if (x == 2) || (x == 3) {
		return 1
	}
	if (x == 1) || (x % 2 == 0) {
		return 0
	}
	for i := 3; i < x; i += 2 {
		if x % i == 0 {
			return 0
		}
	}
	return 1
}

func Min(nums []int) int {
	min := nums[0]
	for _, v := range nums {
		if v < min {
			min = v
		}
	}
	return min
}

func main() {
	var M, N, sum int
	fmt.Scanln(&M)
	fmt.Scanln(&N)
	var prime []int
	for i := M; i <= N; i++ {
		if isPrime(i) == 1 {
			sum += i
			prime = append(prime, i)
		}
	}
	if sum == 0 {
		fmt.Println(-1)
		return
	}
	fmt.Println(sum)
	fmt.Println(Min(prime))

	
}