package main

import "fmt"

type tree struct {
	value int
	left  *tree
	right *tree
}

func main() {
	t := &tree{value: 1,
		left: &tree{value: 2,
			left: &tree{value: 3,
				right: &tree{value: 4,
					left:  &tree{value: 5},
					right: &tree{value: 6},
				},
			},
		},
		right: &tree{value: 2},
	}
	fmt.Println(elWithMaxHigh(t))
}

func elWithMaxHigh(t *tree) [2]int {
	if t == nil {
		return [2]int{0, 0}
	}
	if t.left == nil && t.right == nil {
		return [2]int{t.value, 1}
	}
	l, r := elWithMaxHigh(t.left), elWithMaxHigh(t.right)

	if l[1] > r[1] || l[1] == r[1] && l[0] > r[0] {
		return [2]int{l[0], l[1] + 1}
	}
	return [2]int{r[0], r[1] + 1}
}
