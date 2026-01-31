package main

import "fmt"

func digitPermutation(arr []int) [][]int {
	maskMap := map[[9]int][]int{}
	n := len(arr)
	for i := 0; i < n; i++ {
		mask := [9]int{}
		copyI := arr[i]
		for copyI != 0 {
			if copyI%10 != 0 {
				mask[copyI%10-1]++
			}
			copyI /= 10
		}
		maskMap[mask] = append(maskMap[mask], arr[i])
	}

	result := make([][]int, len(maskMap))
	i := 0
	for _, v := range maskMap {
		result[i] = v
		i++
	}
	return result
}

func main() {
	arr := []int{1230, 99, 23001, 123, 111, 300021, 101010, 90000009, 9}
	fmt.Println(digitPermutation(arr))
}
