package main

import "fmt"

type CMap struct {
	data         map[int]int
	defaultValue int
}

func (cm *CMap) Get(key int) int {
	if cm.data[key] != 0 {
		return cm.data[key]
	}
	return cm.defaultValue
}

func (cm *CMap) Set(key int, value int) {
	cm.data[key] = value
}

func SumSomeElements(arr []int, k int) (int, int) {
	st := CMap{make(map[int]int), -2}
	st.Set(0, -1)
	sum := 0
	n := len(arr)
	for i := 0; i < n; i++ {
		sum += arr[i]
		if st.Get(sum-k) != -2 {
			return st.Get(sum-k) + 1, i
		}
		st.Set(sum, i)
	}

	fmt.Println(st)
	return -1, -1
}

func main() {
	arr := []int{-1, 0, 0, 1, 2, 3, 1, 5, 2, 6, 1, 2, 0, -1, -1, -1}
	k := 4
	fmt.Println(SumSomeElements(arr, k))
}
