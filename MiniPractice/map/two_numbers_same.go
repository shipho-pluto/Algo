package main

import "fmt"

func AreTheSame(a, b int) bool {
	mask1, mask2 := [10]int{}, [10]int{}
	for ; b != 0 && a != 0; func() { a /= 10; b /= 10 }() {
		mask1[a%10]++
		mask2[b%10]++
	}
	for ; b != 0; b /= 10 {
		mask2[b%10]++
	}
	for ; a != 0; a /= 10 {
		mask1[a%10]++
	}

	for i := 0; i < 9; i++ {
		if mask1[i] != mask2[i] {
			return false
		}
	}
	return true
}

func main() {
	n, m := 1203, 321
	fmt.Println(AreTheSame(n, m))
}
