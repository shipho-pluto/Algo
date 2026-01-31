package main

import "fmt"

type tre struct {
	value int
	left  *tre
	right *tre
}

func height(t *tre) int {
	if t == nil {
		return 0
	}
	return max(height(t.left), height(t.right)) + 1
}

func checkIfTreeAVL(t *tre) bool {
	if t != nil {
		return (height(t.left) == height(t.right) ||
			height(t.left) == height(t.right)-1 ||
			height(t.left) == height(t.right)+1) &&
			t.right != nil && t.left != nil
	}
	return true
}

func main() {
	t := &tre{5,
		nil,
		&tre{6,
			nil,
			&tre{7,
				nil,
				nil,
			},
		},
	}

	fmt.Println(checkIfTreeAVL(t))
}
