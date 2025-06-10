package main

import "fmt"

type CArray struct {
	data         []int
	defaultValue int
}

func (c *CArray) Get(index int) int {
	if c.data[index] != 0 {
		return c.data[index]
	}
	return c.defaultValue
}

func (c *CArray) Set(index int, value int) {
	c.data[index] = value
}

func main() {
	{ // max d = 1
		var n int
		max_num := 0
		fmt.Scanln(&n)
		res := 0
		for i := 0; i < n; i++ {
			var el int
			fmt.Scanln(&el)
			res = max(res, max_num*el)
			if el%11 == 0 {
				max_num = max(max_num, el)
			}
		}
		fmt.Println(res)
	}

	{ // min d > 1
		var n int
		fmt.Scan(&n)
		k := CArray{data: make([]int, n), defaultValue: 1000000000000}
		res := 1000000000000
		for i := 0; i < n; i++ {
			var el int
			fmt.Scan(&el)
			for j := 0; j < 9; j++ {
				if el*j%9 == 0 {
					res = min(res, el*k.Get(j))
				}
			}
			k.Set(el%9, min(k.Get(el%9), el))
		}
		fmt.Println(res)
	}
}
