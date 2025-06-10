package main

import "fmt"

func main() {
	// more than n and sum % m

	var n int
	res := 0
	fmt.Scan(&n)
	k := make([]int, 80)
	k50 := make([]int, 80)
	for i := 0; i < n; i++ {
		var x int
		fmt.Scan(&x)
		if x > 50000 {
			res += k[80-x%80]
			k50[x%80]++
		} else {
			res += k50[80-x%80]
		}
		k[x%80]++
	}
	fmt.Println(res)
}
