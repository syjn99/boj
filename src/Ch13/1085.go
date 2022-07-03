package main

import "fmt"

func Min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	var x, y, w, h int
	fmt.Scan(&x)
	fmt.Scan(&y)
	fmt.Scan(&w)
	fmt.Scan(&h)
	// x, y, w-x, h-y 비교
	wMx := w - x
	hMy := h - y

	fmt.Println(Min(Min(x, y), Min(wMx, hMy)))

}