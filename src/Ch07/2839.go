package main

import "fmt"

func Count(kg int) int {
	if kg % 5 == 0 {
		return kg / 5
	}
	var kgPrime int
	for i := kg / 5; i >= 0; i-- {
		kgPrime = kg
		kgPrime -= 5 * i
		if kgPrime % 3 == 0 {
			return i + (kgPrime / 3)
		}
	}
	if kg % 3 == 0 {
		return kg / 3
	}
	return -1
}

func main() {
	var kg int
	fmt.Scanln(&kg)
	fmt.Println(Count(kg))
}