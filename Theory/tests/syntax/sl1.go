package main

import "fmt"

func main() {
	arr := []int{0, 0, 0, 0, 0}
	sl1 := arr[:5]
	sl0 := sl1
	sl0[0] = -1
	sl := make([]int, 5, 6) // pt 0x1A, len = 5, cap = 6
	sl2 := append(sl, 1)    // pt 0x1A, len = 6, cap = 6
	sl2[0] = 1
	sl3 := append(sl2, 2) // pt 0x5B, len = 7, cap = 7
	sl3[0] = 2

	fmt.Println(sl3) // {2, 0, 0, 0, 0, 1, 2}
	fmt.Println(sl2) // {1, 0, 0, 0, 0, 1}
	fmt.Println(sl)  // {1, 0, 0, 0, 0}
	fmt.Println(sl1) // {-1, 0, 0, 0, 0}
	fmt.Println(sl0) // {-1, 0, 0, 0, 0}
	fmt.Println(arr) // {-1, 0, 0, 0, 0}

	st := make(map[int]int)
	stPtr := &st
	fmt.Println(stPtr)

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("panic")
		}
	}()

	var m map[int]int
	v := m[0]
	m[0] = v + 1
	delete(m, 0)
}
