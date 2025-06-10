package main

import (
	"fmt"
	"sort"
)

func sum9(arr []int) int {
	sum := 0
	for _, x := range arr {
		sum += x
	}
	return sum
}

func Hotel(people [][]int) int {
	maxEnd := 0
	for i := 0; i < len(people); i++ {
		maxEnd = max(maxEnd, people[i][1])
	}
	st := make([]int, maxEnd+1)
	for i := 0; i < len(people); i++ {
		st[people[i][0]] += 1
		st[people[i][1]] -= 1
	}

	ans, cnt := 0, 0
	for i := 0; i < len(st); i++ {
		cnt += st[i]
		ans = max(ans, cnt)
	}
	return ans
}

func HotelHard(people [][]int) int {
	var st [][]int
	for i := 0; i < len(people); i++ {
		st = append(st, []int{people[i][0], 1})
		st = append(st, []int{people[i][1], -1})
	}
	sort.Slice(st, func(i, j int) bool {
		if st[i][0] == st[j][0] {
			return st[i][1] < st[j][1]
		}
		return st[i][0] < st[j][0]
	})

	ans, cnt := 0, 0
	for i := 0; i < len(st); i++ {
		cnt += st[i][1]
		ans = max(ans, cnt)
	}
	return ans
}

func main() {
	people := [][]int{{0, 1}, {1, 2}, {2, 3}, {0, 4}}
	fmt.Println(Hotel(people))
	fmt.Println(HotelHard(people))
}
