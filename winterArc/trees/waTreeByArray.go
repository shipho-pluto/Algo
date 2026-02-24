package main

import (
	t "Algoritm/winterArc/trees/struct"
)

func main() {
	arr := []int{0, 1, 2, 4, 5, 6, 7}
	r := createTreeByArray(arr)
	r.Print()
}

func createTreeByArray(arr []int) *t.TreeInt {
	tree := &t.TreeInt{Value: arr[0]}
	startTree(0, len(arr)-1, tree, arr)
	return tree
}

func startTree(l, r int, root *t.TreeInt, arr []int) {
	m := (l + r) / 2
	root.Value = arr[m]
	if l < m {
		root.Left = &t.TreeInt{}
		startTree(l, m-1, root.Left, arr)
	}
	if r > m {
		root.Right = &t.TreeInt{}
		startTree(m+1, r, root.Right, arr)
	}
}
