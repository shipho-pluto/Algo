package main

import "fmt"

type tree struct {
	val   int
	left  *tree
	right *tree
}

func checkSymmetricTree(root1, root2 *tree) bool {
	if root1 == nil && root2 == nil {
		return true
	}
	if root1 == nil || root2 == nil {
		return false
	}
	return root1.val == root2.val && checkSymmetricTree(root1.left, root2.right) && checkSymmetricTree(root1.right, root2.left)
}

func main() {
	t := &tree{1,
		&tree{2,
			&tree{3,
				nil,
				nil},
			&tree{4,
				nil,
				nil},
		},
		&tree{2,
			&tree{4,
				nil,
				nil},
			&tree{3,
				nil,
				nil},
		},
	}
	fmt.Println(checkSymmetricTree(t.left, t.right))
}
