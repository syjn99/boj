package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func Min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func MidArea(heights []int, start, mid, end int) int {
	height, maxArea := heights[mid], heights[mid]
	left, right := mid, mid
	for (start < left) && (end > right){
		if heights[left-1] < heights[right+1] {
			right++
			height = Min(heights[right], height)
		} else {
			left--
			height = Min(heights[left], height)
		}
		maxArea = Max(maxArea, height * (right - left + 1))
	}

	for (right < end) {
		right++
		height = Min(heights[right], height)
		maxArea = Max(maxArea, height * (right - left + 1))
	}

	for (left > start) {
		left--
		height = Min(heights[left], height)
		maxArea = Max(maxArea, height * (right - left + 1))
	}
	return maxArea
}

func DivideConquer(heights []int, start, end int) (area int) {
	if start >= end {
		return heights[start]
	}

	mid := (start + end) / 2

	leftArea := DivideConquer(heights, start, mid)
	rightArea := DivideConquer(heights, mid+1, end)

	midArea := MidArea(heights, start, mid, end)

	area = Max(Max(leftArea, rightArea), midArea)
	return
}


// func DivideConquer(heights []int, start, end int) (area, height, touch int) {
// 	if start >= end {
// 		return heights[start], heights[start], 3
// 	}

// 	mid := (start + end) / 2

// 	leftArea, leftHeight, leftTouch := DivideConquer(heights, start, mid)
// 	rightArea, rightHeight, rightTouch := DivideConquer(heights, mid+1, end)

// 	if leftTouch == 2 || leftTouch == 3 {
// 		leftTouch -= 2
// 		for i := mid+1; i <= end; i++ {
// 			if heights[i] >= leftHeight {
// 				leftArea += leftHeight
// 				if i == end {
// 					leftTouch += 2
// 				}
// 			} else {
// 				break
// 			}
// 		}
// 	}

// 	if rightTouch == 1 || rightTouch == 3 {
// 		rightTouch -= 1
// 		for i := mid; i >= start; i-- {
// 			if heights[i] >= rightHeight {
// 				rightArea += rightHeight
// 				if i == start {
// 					rightTouch += 1
// 				}
// 			} else {
// 				break
// 			}
// 		}
// 	}

// 	if leftArea >= rightArea {
// 		area, height, touch = leftArea, leftHeight, leftTouch
// 	} else {
// 		area, height, touch = rightArea, rightHeight, rightTouch
// 	}
	
// 	if end - start == 1 {
// 		return
// 	}

// 	midLeft, midRight := heights[mid], heights[mid+1]
// 	tempArea, tempHeight, tempTouch := 0, 0, 0
// 	if midLeft > midRight { // 이런 경우에, 오른쪽이 더 작으므로 midRight부터 왼쪽으로 가면서 area 측정
// 		tempArea, tempHeight = midRight, midRight
// 		for i := mid; i >= start; i-- {
// 			if heights[i] >= tempHeight {
// 				tempArea += tempHeight
// 				if i == start {
// 					tempTouch = 1
// 				}
// 			} else {
// 				break
// 			}
// 		}
// 		for i := mid+2; i <= end; i++ {
// 			if heights[i] >= tempHeight {
// 				tempArea += tempHeight
// 				if i == end {
// 					if tempTouch == 0 {
// 						tempTouch = 2
// 					} else {
// 						tempTouch = 3
// 					}
// 				}
// 			} else {
// 				break
// 			}
// 		}
// 	} else {
// 		tempArea, tempHeight = midLeft, midLeft
// 		for i := (mid + 1); i <= end; i++ {
// 			if heights[i] >= tempHeight {
// 				tempArea += tempHeight
// 				if i == end {
// 					tempTouch = 2
// 				}
// 			} else {
// 				break
// 			}
// 		}
// 		for i := mid-1; i >= start; i-- {
// 			if heights[i] >= tempHeight {
// 				tempArea += tempHeight
// 				if i == start {
// 					if tempTouch == 0 {
// 						tempTouch = 1
// 					} else {
// 						tempTouch = 3
// 					}
// 				}
// 			} else {
// 				break
// 			}
// 		}
// 	}

// 	if area >= tempArea {
// 		return
// 	} else {
// 		area, height, touch = tempArea, tempHeight, tempTouch
// 		return
// 	}
// }

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	for {
		S, _ := r.ReadString('\n')
		S = strings.TrimSpace(S)
		testcase := strings.Split(S, " ")
		if testcase[0] == "0" {
			break
		}
		len, _ := strconv.Atoi(testcase[0])
		var heights []int
		for i := 1; i <= len; i++ {
			height, _ := strconv.Atoi(testcase[i])
			heights = append(heights, height)
		}
		maxArea := DivideConquer(heights, 0, len-1)
		fmt.Fprintln(w, maxArea)
	}
}