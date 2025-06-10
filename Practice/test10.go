package main

import (
	"fmt"
)

type CMap10 struct {
	data         map[int]int
	defaultValue int
}

func (cm *CMap10) Get(key int) int {
	if value, exists := cm.data[key]; exists {
		return value
	}
	return cm.defaultValue
}

func (cm *CMap10) Set(key int, value int) {
	cm.data[key] = value
}

func Pref10(arr []int, k int) [2]int {
	n := len(arr)
	table := CMap10{make(map[int]int), -1}
	table.Set(0, 0)
	sum := 0
	for i := 0; i < n; i++ {
		sum += arr[i]
		if table.Get(sum-k) != -1 {
			return [2]int{table.Get(sum - k), i + 1}
		}
		table.Set(sum, i)
	}

	return [2]int{-1, -1}
}

func main() {
	arr := []int{1, -1, 0, 0, 2, 3, 0}
	k := 3
	fmt.Println(Pref10(arr, k))
}
