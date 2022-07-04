package main

import "fmt"

func Min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func isCoprime(x, y int) (int, bool) {
	N := Min(x, y)
	for i := 2; i <= N; i++ {
		if x % i == 0 && y % i == 0 {
			return i, false
		}
	}
	return 0, true
}

func GCD(x, y int) int {
	gcd := 1
	for {
		num, ok := isCoprime(x, y)
		if ok == true {
			break
		}
		gcd *= num
		x /= num
		y /= num
	}
	return gcd
}

func main() {
	var N, K int
	fmt.Scanln(&N, &K)
	if K >= N / 2 {
		K = N - K
	}

	arr := make([]int, K)

	for i := 0; i < K; i++ {
		arr[i] = N
		N--
	}


	for j := K; j >= 1; j-- {
		div := j
		for i := 0; (i < K) && (div != 1); i++ {
			cd := GCD(arr[i], div)
			if cd != 1 {
				arr[i] = arr[i] /  cd
				div = div / cd
			}
		}
	}

	result := 1
	var results []int

	for i := 0; i < K; i++ {
		result *= arr[i]
		if result > 1000000000000000 {
			results = append(results, result % 10007)
			result = 1
		}
	}
	results = append(results, result % 10007)

	var results2 []int
	result = 1


	for i := 0; i < len(results); i++ {
		result *= results[i]
		if result > 1000000000000000 {
			results2 = append(results2, result % 10007)
			result = 1
		}
	}
	results2 = append(results2, result % 10007)

	var results3 []int
	result = 1


	for i := 0; i < len(results2); i++ {
		result *= results2[i]
		if result > 1000000000000000 {
			results3 = append(results3, result % 10007)
			result = 1
		}
	}
	results3 = append(results3, result % 10007)


	fmt.Println(results3[0])
}