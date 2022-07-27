package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type Heap struct {
	arr []int
	size int
}

type Set struct {
	maxheap, minheap *Heap
	size int
}

func (heap *Heap) Swap(i, j int) {
	temp := heap.arr[i]
	heap.arr[i] = heap.arr[j]
	heap.arr[j] = temp 
}

func (heap *Heap) HeapInsert(x int, mode string) {
	heap.arr = append(heap.arr, x)
	heap.size++
	target, targetInd := heap.arr[heap.size], heap.size
	for (targetInd / 2 != 0) {
		parent := heap.arr[targetInd / 2]
		if mode == "min" {
			if parent < target {
				break
			}
		} else {
			if parent > target {
				break
			}
		}
		heap.Swap(targetInd / 2, targetInd)
		targetInd = targetInd / 2
	}
}

func (heap *Heap) HeapDelete(mode string) int {
	if heap.size == 0 {
		return 0
	}
	elem := heap.arr[1]
	heap.Swap(1, heap.size)
	target, targetInd := heap.arr[1], 1
	for ((targetInd*2 + 1) <= heap.size) {
		var leftChild, rightChild int
		if mode == "min" {
			leftChild, rightChild = 0, math.MaxInt
		} else {
			leftChild, rightChild = 0, math.MinInt
		}

		if (targetInd*2 + 1) == heap.size {
			leftChild = heap.arr[targetInd*2]
		} else {
			leftChild, rightChild = heap.arr[targetInd*2], heap.arr[targetInd*2 + 1]
		}

		if mode == "min" {
			if target < leftChild && target < rightChild {
				break
			}
			if leftChild < rightChild {
				heap.Swap(targetInd, targetInd*2)
				targetInd = targetInd*2
			} else {
				heap.Swap(targetInd, targetInd*2 + 1)
				targetInd = targetInd*2 + 1
			}
		} else {
			if target > leftChild && target > rightChild {
				break
			}
			if leftChild > rightChild {
				heap.Swap(targetInd, targetInd*2)
				targetInd = targetInd*2
			} else {
				heap.Swap(targetInd, targetInd*2 + 1)
				targetInd = targetInd*2 + 1
			}
		}
	}
	heap.size--
	heap.arr = heap.arr[:(heap.size+1)]
	return elem
}

func (set *Set) Insert(x int) {
	if set.size == 0 {
		set.maxheap.HeapInsert(x, "max")
		set.size++
		return
	}

	diff := set.maxheap.size - set.minheap.size

	if diff > 0 {
		set.minheap.HeapInsert(x, "min")
	} else {
		set.maxheap.HeapInsert(x, "max")
	}

	if set.maxheap.arr[1] > set.minheap.arr[1] {
		maxTop, minTop := set.maxheap.HeapDelete("max"), set.minheap.HeapDelete("min")
		set.maxheap.HeapInsert(minTop, "max")
		set.minheap.HeapInsert(maxTop, "min")
	}

	// if x > set.maxheap.arr[1] {
	// 	set.minheap.HeapInsert(x, "min")
	// } else {
	// 	set.maxheap.HeapInsert(x, "max")
	// }


	// if diff == 2 {
	// 	elem := set.maxheap.HeapDelete("max")
	// 	set.minheap.HeapInsert(elem, "min")
	// } else if diff == -2 {
	// 	elem := set.minheap.HeapDelete("min")
	// 	set.maxheap.HeapInsert(elem, "max")
	// }
}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	maxheap, minheap := Heap{[]int{0}, 0}, Heap{[]int{0}, 0}
	set := Set{&maxheap, &minheap, 0}

	var N, x int

	fmt.Fscanln(r, &N)

	for i := 0; i < N; i++ {
		fmt.Fscanln(r, &x)
		set.Insert(x)
		fmt.Fprintln(w, set.maxheap.arr[1])
	}
}