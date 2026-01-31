package main

import "fmt"

func InsertFunc(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		elem := arr[i]
		j := i
		for ; j > 0 && arr[j-1] > elem; j-- {
			arr[j] = arr[j-1]
		}
		arr[j] = elem
	}

	return arr
}

func main() {
	arr := []int{123, 123, 12, 34453, 4567, 89, 56, 45, 345, 897, 67, 234}
	fmt.Println(InsertFunc(arr))
}
