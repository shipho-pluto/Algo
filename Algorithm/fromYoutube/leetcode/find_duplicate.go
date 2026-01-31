package main

import "fmt"

func main() {
	arr := []int{1, 3, 4, 2, 2}
	fmt.Println(findDuplicate(arr))
}

func findDuplicate(arr []int) any {
	st := make(map[int]bool)
	for i := range arr {
		if st[arr[i]] {
			return arr[i]
		}
		st[arr[i]] = true
	}
	return -1
}
