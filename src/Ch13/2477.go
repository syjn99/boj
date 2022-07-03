// 육각형에서 작은 부분 넓이 구하는데에서 어려움을 느낌. 방향을 들어오는 순서대로 슬라이스에 저장하고, 패턴을 찾음. 예, 3131 이런 식으로 연속 되면 가운데 있는 1,3에 움직인 거리가 직사각형의 변의 길이와 같음.
package main

import "fmt"

func MaxMin(arr []int) (max int, min int) {
	max, min = arr[0], arr[0]
	for _, v := range arr {
		if max < v {
			max = v
		}
		if min > v {
			min = v
		}
	}
	return
}

func main() {
	var K, D, M int
	fmt.Scanln(&K)

	// 방향마다 얼마만큼 갔는지 기록.
	dirToMeter := make(map[int][]int)
	// 가는 순서 저장하는 배열
	dirArr := make([]int, 12)

	for i := 0; i < 6; i++ {
		fmt.Scanf("%d %d\n", &D, &M)
		dirArr[i], dirArr[i + 6] = D, D

		dirToMeter[D] = append(dirToMeter[D], M)
	}

	area := 1

	for _, v := range dirToMeter {
		if len(v) == 1 {
			area *= v[0]
		}
	}

	cnt := 0
	var index1, index2 int

	for i := 0; i < 12; i++ {
		if cnt == 4 {
			index1 = i - 3
			index2 = i - 2
			break
		}
		length := len(dirToMeter[dirArr[i]])
		if length == 2 {
			cnt++
		} else {
			cnt = 0
		}
	}
	if index1 <= 3 {
		area -= (dirToMeter[dirArr[index1]][0] * dirToMeter[dirArr[index2]][1])
	} else if index1 == 4 {
		area -= (dirToMeter[dirArr[index1]][1] * dirToMeter[dirArr[index2]][1])
	}	else if index1 == 5 {
		area -= (dirToMeter[dirArr[index1]][1] * dirToMeter[dirArr[index2]][0])
	} else {
		area -= (dirToMeter[dirArr[index1]][0] * dirToMeter[dirArr[index2]][0])
	}

	fmt.Println(K * area)
}