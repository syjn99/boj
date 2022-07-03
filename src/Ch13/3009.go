package main

import "fmt"

func main() {
	var x, y int

	xMap := make(map[int]int)
	yMap := make(map[int]int)


	for i := 0; i < 3; i++ {
		fmt.Scan(&x)
		fmt.Scan(&y)
		xMap[x]++
		yMap[y]++
	}

	var resultX, resultY int

	for k := range xMap {
		if xMap[k] == 1 {
			resultX = k
		}
	}

	for k := range yMap {
		if yMap[k] == 1 {
			resultY = k
		}
	}
	
	fmt.Println(resultX, resultY)
}