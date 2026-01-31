package main

import (
	t "Algoritm/Algorithm/methods/struct/Structs"
	"fmt"
)

func main() {
	arr := []int{0, 1, 2, 4, 5, 6, 7}
	root := &t.Tree{}
	createTreeByArray(0, len(arr)-1, arr, root)
	fmt.Println(root.Print())
}

func createTreeByArray(l, r int, arr []int, root *t.Tree) {
	m := (l + r) / 2
	root.Value = arr[m]
	if l < m {
		root.Left = &t.Tree{}
		createTreeByArray(l, m-1, arr, root.Left)
	}
	if r > m {
		root.Right = &t.Tree{}
		createTreeByArray(m+1, r, arr, root.Right)
	}
}
