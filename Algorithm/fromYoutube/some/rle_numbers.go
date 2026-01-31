package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	arr := []int{-10, -9, -8, -3, -2, -1, 1, 4, 5, 2, 3, 9, 8, 11, 0, 12, 13}
	fmt.Println(rleNumbers(arr))
}

func rleNumbers(arr []int) string {
	sort.Ints(arr)
	res := strings.Builder{}
	n := len(arr)
	l, r := 0, 1
	for ; r < n; r++ {
		if arr[r] != arr[r-1]+1 {
			if l != r-1 {
				res.WriteString(strconv.Itoa(arr[l]) + "->" + strconv.Itoa(arr[r-1]) + ", ")
			} else {
				res.WriteString(strconv.Itoa(arr[r-1]) + ", ")
			}
			l = r
		}
	}
	if arr[r-1] == arr[r-2]+1 {
		res.WriteString(strconv.Itoa(arr[l]) + "->" + strconv.Itoa(arr[r-1]))
	} else {
		res.WriteString(strconv.Itoa(arr[r-1]))
	}

	return res.String()
}
