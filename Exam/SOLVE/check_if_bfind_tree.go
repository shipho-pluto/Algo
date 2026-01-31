package main

import "fmt"

type tree struct {
	value int
	left  *tree
	right *tree
}

func minVal(t *tree) int {
	val := t.value
	if t.left != nil {
		val = min(minVal(t.left), val)
	}
	return val
}

func maxVal(t *tree) int {
	val := t.value
	if t.right != nil {
		val = max(maxVal(t.right), val)
	}
	return val
}

func checkTreeForBF(t *tree, minV, maxV int) bool {
	if t != nil {
		if t.value < minV || t.value > maxV {
			return false
		}
		return checkTreeForBF(t.left, minV, t.value) && checkTreeForBF(t.right, t.value, maxV)
	}
	return true
}

func main() {
	t := &tree{5,
		&tree{3,
			&tree{2,
				nil,
				nil,
			},
			nil,
		},
		&tree{6,
			nil,
			&tree{7,
				nil,
				nil,
			},
		},
	}

	fmt.Println(checkTreeForBF(t, minVal(t)-1, maxVal(t)+1))
}
