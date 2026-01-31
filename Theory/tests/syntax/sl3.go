package main

import "fmt"

type acc struct {
	val int
}

func main() {
	s1 := make([]acc, 0, 2) // ptr 0x01 len = 0, cap = 2
	s1 = append(s1, acc{})  // ptr 0x01 len = 1, cap = 2
	s2 := append(s1, acc{}) // ptr 0x01 len = 2, cap = 2

	ac := &s2[0]        // ptr 0x01
	ac.val = 100        // {val = 100}
	fmt.Println(s1, s2) // [acc{100}], [acc{val : 100}, acc{}]

	s1 = append(s2, acc{}) // ptr 0x4D len = 3, cap = 3
	ac.val += 100          // {val = 200}
	fmt.Println(s1, s2)    // [acc{val : 100}, acc{}, acc{}], [acc{val : 200}, acc{}]

	a := make([]int32, 0)
	a = append(a, []int32{1, 2, 3}...) // len = 3, cap = 3
	a = append(a, 4)                   // len = 4, cap = 4
	fmt.Println(len(a), cap(a))

	b := make([]int64, 0)
	b = append(b, []int64{1, 2, 3}...) // len = 3, cap = 3
	b = append(b, 4)                   // len = 4, cap = 6
	fmt.Println(len(b), cap(b))
}
