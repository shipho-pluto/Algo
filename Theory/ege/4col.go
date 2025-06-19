package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	d := make([]int, 500)
	res := 0
	for i := 0; i < n; i++ {
		var el int
		fmt.Scan(&el)
		res += d[500-el%500]
		d[el%500] += 1
	}

	fmt.Println(res)
}
