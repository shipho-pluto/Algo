package main

import "fmt"

func main() {
	// 1 % n == 0, 1 % m != 2 % m

	{ // sum
		var n int
		fmt.Scan(&n)
		res := 0
		k := make([]int, 160)
		k7 := make([]int, 160)
		for i := 0; i < n; i++ {
			var x int
			fmt.Scan(&x)
			if x%7 == 0 {
				for j := 0; j < 160; j++ {
					if x%160 != j {
						res = max(res, x+k[j])
					}
				}
				k7[x%160] = max(k7[x%160], x)
			} else {
				for j := 0; j < 160; j++ {
					if x%160 != j && j%7 == 0 && k7[j] != 0 {
						res = max(res, x+k7[j])
					}
				}
			}
			k[x%160] = max(x, k[x%160])
		}

		fmt.Println(res)
	}

	{ // count
		var n int
		fmt.Scan(&n)
		res := 0
		k := make([]int, 120)
		k9 := make([]int, 120)
		s, s9 := 0, 0
		for i := 0; i < n; i++ {
			var x int
			fmt.Scan(&x)
			if x%9 == 0 {
				res += s - k9[x%120]
				k9[x%9]++
				s9++
			} else {
				res += s9 - k[x%120]
			}
			k[x%120]++
			s++
		}
		fmt.Println(res)
	}
}
