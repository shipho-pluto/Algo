package main

import (
	"fmt"
)

func BuildTriangle(arr [][]int) int {
	st := map[int]int{}
	for i := 0; i < len(arr); i++ {
		st[arr[i][0]] = max(st[arr[i][0]], arr[i][1])
	}
	res := 0
	for _, v := range st {
		res += v
	}
	return res

}

func main() {
	arr := [][]int{{4, 5}, {1, 2}, {6, 4}, {9, 1}, {9, 10}, {6, 6}, {1, 1}}
	fmt.Println(BuildTriangle(arr))
}
