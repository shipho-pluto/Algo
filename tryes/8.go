package main

import "fmt"

func main() {
	a, b := []int{1, 4, 5, 5, 6, 9, 10}, []int{2, 3, 5, 6, 7}
	fmt.Println(multiplaySort(a, b))
	fmt.Println(insertSort(a, b))
	fmt.Println(dInsertSort(a, b))
	fmt.Println(setSort(a, b))

	fmt.Println("-----------------")
	a, b = []int{10, 10, 4, 6, 4, 2, 19, 3}, []int{19, 1, 1, 4, 5, 6}
	fmt.Println(multiplay(a, b))
	fmt.Println(insert(a, b))
	fmt.Println(dInsert(a, b))
	fmt.Println(set(a, b))

	fmt.Println("-----------------")
	a, b = []int{10, 10, 4, 6, 4, 2}, []int{19, 1, 1, 4, 5, 6}
	fmt.Println(infoEqInArrays(a, b))

	fmt.Println("-----------------")
	a1, a2 := [][]int{{0, 2}, {5, 10}, {13, 23}, {24, 25}}, [][]int{{1, 5}, {8, 12}, {15, 24}, {25, 26}}
	fmt.Println(subIntervals(a1, a2))
}

func subIntervals(a1 [][]int, a2 [][]int) [][]int {
	var res [][]int
	for l, r := 0, 0; l < len(a1) && r < len(a2); {
		maxStart := max(a1[l][0], a2[r][0])
		minEnd := min(a1[l][1], a2[r][1])

		if maxStart < minEnd {
			res = append(res, []int{maxStart, minEnd})
		}

		if a1[l][1] < a2[r][1] {
			l++
		} else {
			r++
		}
	}

	return res
}

func infoEqInArrays(a []int, b []int) []int {
	res := make([]int, 0, len(a))
	c := 0
	st1, st2 := make(map[int]bool, len(a)), make(map[int]bool, len(b))
	for i := range a {
		if !st1[a[i]] {
			st1[a[i]] = true
			if st2[a[i]] {
				c++
			}
		}

		if !st2[b[i]] {
			st2[b[i]] = true
			if st1[b[i]] {
				c++
			}
		}

		res = append(res, c)
	}

	return res
}

func set(a, b []int) []int {
	var res []int

	st := make(map[int]bool)
	for i := range a {
		st[a[i]] = true
	}
	for i := range b {
		st[b[i]] = true
	}

	for i := range st {
		if st[i] {
			res = append(res, i)
		}
	}

	return res
}

func setSort(a, b []int) []int {
	var res []int
	cur := min(a[0], b[0]) - 1
	for l, r := 0, 0; l < len(a) || r < len(b); {
		if l < len(a) && r < len(b) {
			if a[l] < b[r] {
				if a[l] != cur {
					res = append(res, a[l])
				}
				l++
			} else if a[l] > b[r] {
				if b[r] != cur {
					res = append(res, b[r])
				}
				r++
			} else {
				res = append(res, b[r])
				cur = b[r]
				l, r = l+1, r+1
			}
		} else if l < len(a) {
			res = append(res, a[l])
			l++
		} else {
			res = append(res, b[r])
			r++
		}
	}

	return res
}

func dInsert(a, b []int) []int {
	var res []int
	st := make(map[int]int)
	for i := range a {
		st[a[i]]++
	}

	for i := range b {
		if st[b[i]] == 0 {
			res = append(res, b[i])
		} else {
			st[b[i]] -= st[b[i]] + 1
		}
	}

	for i := range st {
		if st[i] > 0 {
			for range st[i] {
				res = append(res, i)
			}
		}
	}

	return res
}

func dInsertSort(a, b []int) any {
	var res []int
	cur := min(a[0], b[0]) - 1
	for l, r := 0, 0; l < len(a) || r < len(b); {
		if l < len(a) && r < len(b) {
			if a[l] < b[r] {
				if a[l] != cur {
					res = append(res, a[l])
				}
				l++
			} else if a[l] > b[r] {
				if b[r] != cur {
					res = append(res, b[r])
				}
				r++
			} else {
				cur = a[l]
				l, r = l+1, r+1
			}
		} else if l < len(a) {
			res = append(res, a[l])
			l++
		} else {
			res = append(res, b[r])
			r++
		}
	}

	return res
}

func insert(a, b []int) []int {
	var res []int
	st := make(map[int]bool, len(b))
	for i := range b {
		st[b[i]] = true
	}

	for i := range a {
		if !st[a[i]] {
			res = append(res, a[i])
		}
	}
	return res
}

func insertSort(a, b []int) []int {
	var res []int
	cur := min(a[0], b[0]) - 1
	for l, r := 0, 0; l < len(a); {
		if r < len(b) && a[l] < b[r] || r == len(b) {
			if a[l] != cur {
				res = append(res, a[l])
			}
			l++
		} else if r < len(b) && a[l] > b[r] {
			r++
		} else {
			cur = a[l]
			l, r = l+1, r+1
		}
	}
	return res
}

func multiplay(a, b []int) []int {
	var res []int
	st := make(map[int]int, len(b))
	for i := range b {
		st[b[i]]++
	}

	for i := range a {
		if st[a[i]] != 0 {
			res = append(res, a[i])
			st[a[i]]--
		} else {
			delete(st, a[i])
		}
	}
	return res
}

func multiplaySort(a, b []int) []int {
	var res []int
	cur := min(a[0], b[0]) - 1
	for l, r := 0, 0; l < len(a) && r < len(b); {
		if a[l] < b[r] {
			l++
		} else if a[l] == b[r] && a[l] != cur {
			res = append(res, a[l])
			cur = a[l]
			l, r = l+1, r+1
		} else {
			r++
		}
	}

	return res
}
