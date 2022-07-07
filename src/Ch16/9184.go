package main

import (
	"bufio"
	"fmt"
	"os"
)

type Coordinate struct {
	x, y, z int
}

var Values map[Coordinate]int = make(map[Coordinate]int)
var isValue map[Coordinate]bool = make(map[Coordinate]bool)

func calValue(coo Coordinate) int {
	x, y, z := coo.x, coo.y, coo.z

	if x <= 0 || y <= 0 || z <= 0 {
		isValue[coo] = true
		Values[coo] = 1
		return Values[coo]
	}

	if x < y && y < z {
		return wCal(Coordinate{x, y, z - 1}) + wCal(Coordinate{x, y - 1, z - 1}) - wCal(Coordinate{x, y - 1, z})
	} else {
		return wCal(Coordinate{x - 1, y, z}) + wCal(Coordinate{x - 1, y - 1, z}) + wCal(Coordinate{x - 1, y, z - 1}) - wCal(Coordinate{x - 1, y - 1, z - 1})
	}
}

func wCal(coo Coordinate) int {
	if isValue[coo] {
		return Values[coo]
	}


	x, y, z := coo.x, coo.y, coo.z
	if x <= 0 || y <= 0 || z <= 0 {
		isValue[coo] = true
		Values[coo] = 1
		return Values[coo]
	}

	if x > 20 || y > 20 || z > 20 {
		return wCal(Coordinate{20, 20, 20})
	}

	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			for k := 0; k < z; k++ {
				tempCoo1 := Coordinate{i, j, k + 1}
				Values[tempCoo1] = calValue(tempCoo1)
				isValue[tempCoo1] = true
			}
			tempCoo2 := Coordinate{i, j + 1, z}
			Values[tempCoo2] = calValue(tempCoo2)
			isValue[tempCoo2] = true
		}
		tempCoo3 := Coordinate{i + 1, y, z}
		Values[tempCoo3] = calValue(tempCoo3)
		isValue[tempCoo3] = true
	}
	return Values[coo]
}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	var x, y, z int


	for {
		fmt.Fscanln(r, &x, &y, &z)
		if x == -1 && y == -1 && z == -1 {
			break
		}
		coo := Coordinate{x, y, z}
		fmt.Fprintf(w, "w(%d, %d, %d) = %d\n", x, y, z, wCal(coo))
	}
	w.Flush()
}