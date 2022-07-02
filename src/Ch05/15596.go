package main

func sum(a []int) int {
	var result int = 0
    for _, x := range a {
		result = result + x
	}
	return result
}
