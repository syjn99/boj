package main

import "fmt"

func main() {
 var N, M, x, sum int
 fmt.Scan(&N)
 fmt.Scan(&M)

 arr := make([]int, N)
 for i := 0; i < N; i++ {
	fmt.Scan(&x)
	arr[i] = x
 }

 sum = 0

 for i := 0; i < N - 2; i++ {
	for j := i + 1; j < N - 1; j++ {
		for k := j + 1; k < N; k++ {
			temp := arr[i] + arr[j] + arr[k]
			if sum < temp && temp <= M  {
				sum = temp
			}
		}
	}
 }

 fmt.Println(sum)
}