package main

import (
	"fmt"
)

func main() {
	var a, b, c, d int
	fmt.Scanln(&a, &b, &c, &d)

	if c <= a/10 {
		fmt.Println(d * c)
	} else {
		res := a
		d -= b
		for d > b {
			d -= b
			res += a
		}
		if d > 0 {
			res += d * c
		}
		fmt.Println(res)
	}
}
