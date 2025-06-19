package main

import "fmt"

func TrueTurtle(arr [][]int) int {
	cnt := 0
	ans := make(map[int]bool)
	n := len(arr)
	for i := 0; i < n; i++ {
		a, b := arr[i][0], arr[i][1]
		if !(a < 0 || b < 0 || a+b != n-1 || ans[a] == true) {
			cnt++
			ans[a] = true
		}
	}
	return cnt
}

func main() {
	arr := [][]int{{3, 0}, {2, 1}, {1, 2}, {1, 3}}
	fmt.Println(TrueTurtle(arr))
}
