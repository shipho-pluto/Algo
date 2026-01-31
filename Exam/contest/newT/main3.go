package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	e, p := 0, 1
	for p < n {
		e += 1
		p *= 2
	}
	fmt.Println(e)
}
