package main

import (
	"fmt"
	"os"
)

func main() {
	var n int
	res := map[int]int{}
	maxSum := 0
	fmt.Fscan(os.Stdin, &n)
	for i := 0; i < n; i++ {
		var k int
		fmt.Fscan(os.Stdin, &k)
		if res[5-(k%5)] != 0 {
			maxSum = max(maxSum, res[5-(k%5)]+k)
		}
		res[k%5] = max(k, res[k%5])
	}
	fmt.Println(maxSum)
}
