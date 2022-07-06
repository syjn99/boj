package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var results []int
var N int
var powers [][]int


func Min(results []int) int {
	min := results[0]

	for _, v := range results {
		if v < min {
			min = v
		}
	}

	return min
}

func Score(team []int) float64 {
	result := 0.0
	for i := 0; i < len(team); i++ {
		p1 := team[i]
		for j := 0; j < len(team); j++ {
			p2 := team[j]
			result += float64(powers[p1][p2])
		}
	}
	return result
}

func MakeTeam(team1 []int, selected map[int]bool) {
	if len(team1) == N/2 {
		var team2 []int
		m := make(map[int]bool)
		for _, v := range team1 {
			m[v] = true
		}
		for i := 0; i < N; i++ {
			if !m[i] {
				team2 = append(team2, i)
			}
		}
		results = append(results, int(math.Abs(Score(team1) - Score(team2))))
		return
	}

	tempTeam := make([]int, len(team1))
	for i, v := range team1 {
		tempTeam[i] = v
	}

	index := 0

	if len(team1) != 0 {
		index = team1[len(team1) - 1]
	}

	for i := index; i < N; i++ {
		if !selected[i] && len(team1) < N/2 {
			selected[i] = true
			team1 = append(team1, i)
			MakeTeam(team1, selected)
			team1 = team1[:(len(team1) - 1)]
			selected[i] = false
		}
	}

}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var x int

	fmt.Fscanln(r, &N)
	powers = make([][]int, N)
	for i := 0; i < N; i++ {
		powers[i] = make([]int, N)
	}

	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			fmt.Fscan(r, &x)
			powers[i][j] = x
		}
	}

	var team1 []int
	selected := make(map[int]bool)

	MakeTeam(team1, selected)
	
	fmt.Fprintln(w, Min(results))
}