package main

import "fmt"

type TreeIn3 struct {
	left  *TreeIn3
	right *TreeIn3
	value int
}

func Abs3(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Height(t *TreeIn3, loop int) int {
	if t != nil {
		loop = max(Height(t.left, loop+1), Height(t.right, loop+1))
	}
	return loop
}

func IsTreeAVL(t *TreeIn3) bool {
	if t != nil {
		leftHieght := Height(t.left, 1)
		rightHieght := Height(t.right, 1)
		return Abs3(leftHieght-rightHieght) <= 1 && t.left != nil && nil != t.right
	}
	return true
}

func main() {
	t := TreeIn3{nil, &TreeIn3{&TreeIn3{nil, nil, 25}, nil, 75}, 50}
	fmt.Println(IsTreeAVL(&t))
}
