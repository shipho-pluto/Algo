package main

import (
	"fmt"
	"strconv"
)

type Tree struct {
	val   int
	left  *Tree
	right *Tree
}

func (t *Tree) Print() string {
	return printTree(t)
}

func printTree(t *Tree) string {
	if t != nil {
		return strconv.Itoa(t.val) + " " + printTree(t.left) + " " + printTree(t.right)
	}
	return "end"
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	tree := createByArray(arr)

	fmt.Println(tree.Print())
}

func createByArray(arr []int) *Tree {
	root := &Tree{}
	create(0, len(arr)-1, arr, root)
	return root
}

func create(l int, r int, arr []int, root *Tree) {
	m := (l + r) / 2
	root.val = arr[m]
	if l < m {
		root.left = &Tree{}
		create(l, m-1, arr, root.left)
	}

	if r > m {
		root.right = &Tree{}
		create(m+1, r, arr, root.right)
	}
}
