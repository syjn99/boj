package main

import (
	"bufio"
	"fmt"
	"os"
)

type Node struct {
	value int
	parent, left, right *Node
}

type Tree struct {
	root *Node
	size, leftCnt, rightCnt int
}

func Init() *Tree {
	tree := Tree{}
	tree.root = nil
	tree.size, tree.leftCnt, tree.rightCnt = 0, 0, 0
	return &tree
}

func (tree *Tree) Insert(x int) {
	tree.size++
	node := Node{x, nil, nil, nil}
	if tree.root == nil {
		tree.root = &node
		return
	}

	walk := tree.root
	var prev *Node

	if x < walk.value {
		tree.leftCnt++
		for (walk != nil) && (x < walk.value) {
			prev = walk
			walk = walk.left
		}
		prev.left = &node
		node.parent = prev
		if walk != nil {
			node.left = walk
		}
	} else {
		tree.rightCnt++
		for (walk != nil) && (x >= walk.value) {
			prev = walk
			walk = walk.right
		}
		prev.right = &node
		node.parent = prev
		if walk != nil {
			node.right = walk
		}
	}

	diff := tree.leftCnt - tree.rightCnt
	if diff == 2 || diff == 1 {
		A, B := tree.root, tree.root.left
		A.left = nil
		B.right = A
		tree.root = B
		tree.leftCnt--
		tree.rightCnt++
	} else if diff == -2 {
		A, B := tree.root, tree.root.right
		A.right = nil
		B.left = A
		tree.root = B
		tree.rightCnt--
		tree.leftCnt++
	}
}



func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	tree := Init()

	var N, x int

	fmt.Fscanln(r, &N)

	for i := 0; i < N; i++ {
		fmt.Fscanln(r, &x)
		tree.Insert(x)
		fmt.Fprintln(w, tree.root.value)
	}
}