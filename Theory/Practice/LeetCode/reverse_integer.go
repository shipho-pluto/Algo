package main

import "fmt"

func Reverse(x int) int {
	res := 0
	for ; x != 0; x /= 10 {
		res = 10*res + x%10
	}
	if int(int32(res)) != res {
		res = 0
	}
	return res
}

func main() {
	x := 1534236469
	fmt.Println(Reverse(x))
}
