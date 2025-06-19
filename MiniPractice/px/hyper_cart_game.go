package main

import (
	"fmt"
	"sort"
)

func HowManyCombinationsLessThanK(arr []int, k int) int {
	st := map[int][]int{}
	for i := range arr {
		if st[arr[i]] == nil {
			st[arr[i]] = []int{arr[i], 1}
		} else {
			st[arr[i]][1]++
		}
	}
	var ms [][]int
	for _, v := range st {
		ms = append(ms, v)
	}
	sort.Slice(ms, func(i, j int) bool {
		return ms[i][0] < ms[j][0]
	})

	res := 0
	l, r, m := 1, 1, len(ms)
	pref := make([][]int, m+1)
	pref[0] = []int{0, 0, 0}
	for i := 0; i < m; i++ {
		if ms[i][1] > 2 {
			res++
		}
		if ms[i][1] > 1 {
			pref[i+1] = []int{ms[i][0], pref[i][1], pref[i][2] + 1}
		} else {
			pref[i+1] = []int{ms[i][0], pref[i][1] + 1, pref[i][2]}
		}
	}
	arrK := make([]int, k+1)
	arrK[0] = 0
	for i := range k {
		arrK[i+1] += arrK[i] + (i + 1)
	}
	for l < m+1 {
		for r < m+1 && k*pref[l][0] > pref[r][0] {
			r++
		}
		// Слишком много условий проверять, а так всё верно
		//
		// Новая идея - посчитать сначала первый шаг, а все последующие считать только по приходящим числам,
		// если пара - то одни действия, если одиночка - то другие
		//
		//
		//c2 := pref[r-1][2] - pref[l-1][2]
		//c1 := pref[r-1][1] - pref[l-1][1]
		//if c2 >= 1 && c1 >= 1 || c2 >= 2 {
		//	res += c2*3*(c2-1+c1) + c1*(c2-1)
		//}
		//if c1 >= 3 {
		//	res += arrK[c1-2] * 6
		//}
		//if c1 >= 2 && c2 >= 1 {
		//	res += arrK[c1-1] * 6
		//}
		l++
	}

	return res
}

func main() {
	arr := []int{1, 2, 3, 3, 3}
	k := 5
	fmt.Println(HowManyCombinationsLessThanK(arr, k))
}
