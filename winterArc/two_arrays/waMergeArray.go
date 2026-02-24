package main

import "fmt"

func main() {
	arr1, arr2 := []int{1, 3, 3, 3, 3, 3, 5}, []int{2, 3, 4}
	fmt.Println(merge(arr1, arr2))
	fmt.Println(mergeCount(arr1, arr2))
	fmt.Println(unmerge(arr1, arr2))
	fmt.Println(unmergeCount(arr1, arr2))
	arr1, arr2 = []int{1, 3, 5, 0, 0}, []int{2, 4}
	fmt.Println(mergeInOne(arr1, arr2))

}

func mergeInOne(arr1 []int, arr2 []int) []int {
	for i, l, r := len(arr1)-1, len(arr1)-len(arr2)-1, len(arr2)-1; i > -1; i-- {
		if l > -1 && (r > -1 && arr1[l] > arr2[r] || r == -1) {
			arr1[i] = arr1[l]
			l--
		} else {
			arr1[i] = arr2[r]
			r--
		}
	}
	return arr1
}

func merge(arr1 []int, arr2 []int) []int {
	var res []int
	for l, r := 0, 0; l < len(arr1) || r < len(arr2); {
		if l < len(arr1) && (r < len(arr2) && arr1[l] < arr2[r] || r == len(arr2)) {
			res = append(res, arr1[l])
			l++
		} else {
			res = append(res, arr2[r])
			r++
		}
	}
	return res
}

func mergeCount(arr1 []int, arr2 []int) []int {
	var res []int
	for l, r := 0, 0; l < len(arr1) || r < len(arr2); {
		if l < len(arr1) && (r < len(arr2) && arr1[l] <= arr2[r] || r == len(arr2)) {
			if l < len(arr1) && r < len(arr2) && arr1[l] == arr2[r] {
				r++
			}
			if l > 0 && arr1[l] != arr1[l-1] || l == 0 {
				res = append(res, arr1[l])
			}
			l++
		} else {
			if r > 0 && arr2[r] != arr2[r-1] || r == 0 {
				res = append(res, arr2[r])
			}
			r++
		}
	}
	return res
}

func unmerge(arr1 []int, arr2 []int) []int {
	var res []int
	cur := arr1[0] - 1
	for l, r := 0, 0; l < len(arr1) || r < len(arr2); {
		if l < len(arr1) && r < len(arr2) {
			if arr1[l] < arr2[r] && arr1[l] != cur {
				res = append(res, arr1[l])
				l++
			} else if arr1[l] > arr2[r] && arr2[r] != cur {
				res = append(res, arr2[r])
				r++
			} else if arr1[l] == arr2[r] {
				cur = arr1[l]
				l, r = l+1, r+1
			} else if arr1[l] == cur {
				l++
			} else {
				r++
			}
		} else if l < len(arr1) {
			if arr1[l] != cur {
				res = append(res, arr1[l])
			}
			l++
		} else {
			if arr2[r] != cur {
				res = append(res, arr2[r])
			}
			r++
		}
	}
	return res
}

func unmergeCount(arr1 []int, arr2 []int) []int {
	var res []int
	for l, r := 0, 0; l < len(arr1) || r < len(arr2); {
		if l < len(arr1) && r < len(arr2) {
			if arr1[l] < arr2[r] {
				res = append(res, arr1[l])
				l++
			} else if arr1[l] > arr2[r] {
				res = append(res, arr2[r])
				r++
			} else if arr1[l] == arr2[r] {
				l, r = l+1, r+1
			}
		} else if l < len(arr1) {
			res = append(res, arr1[l])
			l++
		} else {
			res = append(res, arr2[r])
			r++
		}
	}
	return res
}
