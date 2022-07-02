package main

import (
	"fmt"
)

type Person struct {
	x, y int
}

// isBigger : 1 = A > B, 0 = A = B, -1 = A < B

func isBigger(A, B Person) int {
	if (A.x > B.x) && (A.y > B.y) {
		return 1
	} else if (A.x < B.x) && (A.y < B.y) {
		return -1
	}
	return 0
}

func main() {
	var N, x, y int
	fmt.Scanln(&N)
	people := make([]Person, N)
	rank := make(map[int]int)
	for i := 0; i < N; i++ {
		fmt.Scan(&x)
		fmt.Scan(&y)
		people[i] = Person{x, y}
	}

	for person := range people {
		rank[person] = 1
	}

	for i := 0; i < N - 1; i++ {
		for j := i; j < N; j++ {
			A, B := people[i], people[j]
			if isBigger(A, B) == 1 {
				rank[j]++
			} else if isBigger(A, B) == -1 {
				rank[i]++
			}
		}
	}
	for person := range people {
		fmt.Printf("%d ", rank[person])
	}
	fmt.Printf("\n")
}