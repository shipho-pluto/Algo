package main

import t "Algoritm/Algorithm/methods/struct/Structs"

func main() {

}

func isTreeNodeAverage(root *t.Tree, min, max int) bool {
	if root == nil {
		return true
	}
	return root.Value <= max && root.Value >= min &&
		isTreeNodeAverage(root.Left, min, max) &&
		isTreeNodeAverage(root.Right, min, max)
}
