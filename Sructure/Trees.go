package main

import "fmt"

type Tree struct {
	value int
	left  *Tree
	right *Tree
}

func (t *Tree) New(value int, left, right *Tree) {
	t.value = value
	t.left = left
	t.right = right
}

func (t *Tree) recursive() int {
	sum := t.value
	if t.left != nil {
		sum += t.left.recursive()
	}
	if t.right != nil {
		sum += t.right.recursive()
	}
	return sum
}

func main() {
	tree := &Tree{value: 20,
		left: &Tree{value: 7,
			left: &Tree{value: 4,
				left: nil,
				right: &Tree{value: 6,
					left:  nil,
					right: nil,
				},
			},
			right: &Tree{value: 9,
				left:  nil,
				right: nil,
			},
		},
		right: &Tree{value: 35,
			left: &Tree{value: 31,
				left: &Tree{value: 28,
					left:  nil,
					right: nil,
				},
				right: nil,
			},
			right: &Tree{value: 40,
				left: &Tree{value: 38,
					left:  nil,
					right: nil,
				},
				right: &Tree{value: 52,
					left:  nil,
					right: nil,
				},
			},
		},
	}
	fmt.Println(tree.recursive())
}
