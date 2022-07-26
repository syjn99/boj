package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type MinHeap struct {
	arr []int
	size int
}

var heap MinHeap

func (h *MinHeap) Swap(i, j int) {
	temp := h.arr[i]
	h.arr[i] = h.arr[j]
	h.arr[j] = temp 
}

func (h *MinHeap) Delete() int {
	if h.size == 0 {
		return 0
	}
	elem := h.arr[1]
	h.Swap(1, h.size)
	target, targetInd := h.arr[1], 1
	for ((targetInd*2 + 1) <= h.size) {
		leftChild, rightChild := 0, math.MaxInt
		if (targetInd*2 + 1) == h.size {
			leftChild = h.arr[targetInd*2]
		} else {
			leftChild, rightChild = h.arr[targetInd*2], h.arr[targetInd*2 + 1]
		}
		if target < leftChild && target < rightChild {
			break
		}
		if leftChild < rightChild {
			h.Swap(targetInd, targetInd*2)
			targetInd = targetInd*2
		} else {
			h.Swap(targetInd, targetInd*2 + 1)
			targetInd = targetInd*2 + 1
		}
	}
	h.size--
	h.arr = h.arr[:(h.size+1)]
	return elem
}

func (h *MinHeap) Insert(x int) {
	h.arr = append(h.arr, x)
	h.size++
	target, targetInd := h.arr[h.size], h.size
	for (targetInd / 2 != 0) {
		parent := h.arr[targetInd / 2]
		if parent < target {
			break
		}
		h.Swap(targetInd / 2, targetInd)
		targetInd = targetInd / 2
	}
}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var N, x int

	fmt.Fscanln(r, &N)

	heap.arr = []int{0}
	heap.size = 0

	for i := 0; i < N; i++ {
		fmt.Fscanln(r, &x)
		if x == 0 {
			fmt.Fprintln(w, heap.Delete())
		} else {
			heap.Insert(x)
		}
	}
}