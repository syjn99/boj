package main

import "fmt"

func main() {
	var T, H, W, N int
	fmt.Scanln(&T)

	for i := 0; i < T; i++ {
		fmt.Scan(&H)
		fmt.Scan(&W)
		fmt.Scan(&N)
		first := (N / H) + 1
		third := N % H
		if third == 0 {
			third = H
			first -= 1
		}

		if first >= 10 {
			fmt.Printf("%d%d\n", third, first)
		} else {
			fmt.Printf("%d0%d\n", third, first)
		}
		// 10 /6 = 1 ... 4 => 402
		// 5 / 6 = 0 ... 5 => 501
		// 6 / 6 = 1 ... 0 => 601
		// 12 / 6 = 2 ... 0 => 602
	}
}