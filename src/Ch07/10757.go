package main

import (
	"fmt"
	"strconv"
)

func Min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	var A, B string
	fmt.Scan(&A)
	fmt.Scan(&B)
	AComNum := len(A) / 10
	BComNum := len(B) / 10
	length := Max(AComNum, BComNum) + 1
	ComNum := Min(AComNum, BComNum) + 1
	if len(A) <= 18 && len(B) <= 18 {
		SmallA, _ := strconv.Atoi(A)
		SmallB, _ := strconv.Atoi(B)
		fmt.Println(SmallA + SmallB)
		return
	}

	var results []int

	carry := 0

	for i := 0; i < (ComNum * 10) ; i += 10 {
		start := len(A) - i - 10
		if start < 0 {
			start = 0
		}
		APart, _ := strconv.Atoi(A[(start):(len(A) - i)])
		start = len(B) - i - 10
		if start < 0 {
			start = 0
		}
		BPart, _ := strconv.Atoi(B[(start):(len(B) - i)])
		result := APart + BPart + carry
		if result >= 10000000000 {
			result -= 10000000000
			carry = 1
		} else {
			carry = 0
		}
		results = append(results, result)
	}

	for k := (ComNum * 10);; k += 10 {
		if AComNum == (length - 1) {
			start := len(A) - k - 10
			if start < 0 {
				start = 0
			}
			if (len(A) - k) < 0 {
				break
			}
			APart, _ := strconv.Atoi(A[(start):(len(A) - k)])
			APart += carry
			carry = 0
			if APart >= 10000000000 {
				APart -= 10000000000
				carry = 1
			}
			results = append(results, APart)
		} else {
			start := len(B) - k - 10
			if start < 0 {
				start = 0
			}
			if (len(B) - k) < 0 {
				break
			}
			BPart, _ := strconv.Atoi(B[(start):(len(B) - k)])
			BPart += carry
			carry = 0
			if BPart >= 10000000000 {
				BPart -= 10000000000
				carry = 1
			}
			results = append(results, BPart)
		}
	}


	for j := (len(results) - 1); j >= 0; j-- {
		if j == (len(results) - 1) && results[j] == 0 {
			continue
		}
		if j == (len(results) - 1) {
			fmt.Printf("%d", results[j])
			continue
		}
		toString := strconv.Itoa(results[j])
		for l := 0; l < (10 - len(toString)); l++ {
			fmt.Printf("0")
		}
		fmt.Printf("%s", toString)
	}

	fmt.Printf("\n")

}