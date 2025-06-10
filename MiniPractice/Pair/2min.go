package main

import (
	"fmt"
	"os"
)

func Abs3(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	var n int
	d := 17
	fmt.Fscan(os.Stdin, &n)
	m := []int{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1}
	maximum := 0
	for i := 0; i < n; i++ {
		var el int
		fmt.Fscan(os.Stdin, &el)
		if m[el%d] != -1 {
			maximum = max(maximum, Abs3(el-m[el%d]))
			m[el%d] = min(el, m[el%d])
		} else {
			m[el%d] = el
		}
	}
	fmt.Println(maximum)
}
