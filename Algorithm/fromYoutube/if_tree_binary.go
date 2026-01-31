package main

import "fmt"

type tree struct {
	val   int
	left  *tree
	right *tree
}

func (t *tree) isBinaryTree(maxV, minV int) bool {
	if t != nil {
		if t.left != nil && t.left.val < minV || t.right != nil && t.right.val > maxV {
			return false
		}

		return t.right.isBinaryTree(maxV, t.val) && t.left.isBinaryTree(t.val, minV)
	}
	return true
}

func (t *tree) max() int {
	if t.right != nil {
		return t.right.max()
	}
	return t.val
}

func (t *tree) min() int {
	if t.left != nil {
		return t.left.min()
	}
	return t.val
}

func main() {
	t := &tree{val: 0,
		left: &tree{val: -1},
		right: &tree{val: 10,
			right: &tree{val: 11,
				right: &tree{val: 12}},
		},
	}

	fmt.Println(t.isBinaryTree(t.max()+1, t.min()-1))
}
