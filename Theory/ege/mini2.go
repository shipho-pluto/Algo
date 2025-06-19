package main

import "fmt"

func AbsMM(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	// sum % n and multi % m
	{ // d = 1
		var n int
		fmt.Scanln(&n)
		d13 := make([]int, 2)
		ch := make([]int, 2)
		count := 0
		for i := 0; i < n; i++ {
			var x int
			fmt.Scan(&x)
			if x%13 == 0 {
				count += ch[x%2]
				d13[x%2]++
			} else {
				count += d13[x%2]
			}
			ch[x%2]++
		}

		fmt.Println(count)
	}

	{ // d > 1
		var n int
		fmt.Scan(&n)
		k := make([]int, 18)
		res := 0
		for i := 0; i < n; i++ {
			var x int
			fmt.Scan(&x)
			for j := 0; j < 18; j++ {
				if AbsMM(x-j)%2 == 1 && (x*j)%9 == 0 {
					res += k[j]
				}
			}
			k[x%18]++
		}
		fmt.Println(res)
	}
}
