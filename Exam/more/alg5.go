package main

import (
	"Algoritm/Theory/Sructure"
	"fmt"
)

func main() {
	arr := []int{7, 2, 4, 5, 3, 2, 5, 1, 5, 4}
	stack := Struct.Stack{}
	stack.Push([]int{0, -1})
	res := make([]int, len(arr))
	for i, v := range arr {
		if v < stack.Look().([]int)[0] {
			index := i
			for v < stack.Look().([]int)[0] {
				if stack.Look() == nil {
					break
				}
				array := stack.Pop()
				res[array.([]int)[1]] = index
			}
		}
		stack.Push([]int{v, i})
	}
	for i, v := range res {
		if v == 0 {
			res[i] = stack.Look().([]int)[1] + 1
		}
	}
	fmt.Println(res)
}
