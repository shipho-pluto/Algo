package main

import "fmt"

func CounterSort(arr []int, dif int) []int {
	var res []int
	count := make([]int, dif)
	for i := 0; i < len(arr); i++ {
		count[arr[i]]++
	}
	for i, val := range count {
		for j := 0; j < val; j++ {
			res = append(res, i)
		}
	}
	return res
}

type People struct {
	name string
	age  int
}

func CounterSortObj(arr []People, dif int) []People {
	res := make([]People, len(arr))
	count := make([]int, dif)
	for i := 0; i < len(arr); i++ {
		count[arr[i].age]++
	}
	for i := 0; i < dif-1; i++ {
		count[i+1] += count[i]
	}
	for i := 0; i < len(arr); i++ {
		res[count[arr[i].age]-1] = arr[i]
	}
	return res
}

func main() {
	arr := []int{1, 3, 4, 2, 4, 3, 2, 4, 3, 2, 3, 5, 4, 5, 5, 1, 0, 0, 0, 1, 2, 3, 1, 0, 0, 0, 1}
	fmt.Println(CounterSort(arr, 6))

	var arrObj = []People{
		{name: "Nick", age: 23},
		{name: "Jack", age: 22},
		{name: "Brandon", age: 97},
		{name: "Michael", age: 17},
		{name: "Robert", age: 9},
		{name: "Tom", age: 18}}
	fmt.Println(CounterSortObj(arrObj, 100))
}
