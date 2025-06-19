package main

import "fmt"

func moveZeroes1(arr []int) {
	newArr := make([]int, len(arr))
	index := 0
	for i := 0; i < len(newArr); i++ {
		if arr[i] != 0 {
			newArr[index] = arr[i]
			index++
		}
	}
	copy(arr, newArr)
}

func moveZeroes(arr []int) {
	n := len(arr)
	l := n
	for i := 0; i < n; i++ {
		if arr[i] == 0 {
			l = i
			break
		}
	}
	r := l + 1
	for ; r < n; r++ {
		if arr[r] != 0 {
			arr[l], arr[r] = arr[r], arr[l]
			l++
		}
	}
}

func main() {
	arr := []int{0, 1, 0, 3, 12}
	moveZeroes(arr)
	fmt.Println(arr)
}
