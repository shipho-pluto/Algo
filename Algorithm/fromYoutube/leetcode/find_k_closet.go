package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println(findClosestElements(arr, 4, 100))
}

func findClosestElements(arr []int, k int, x int) []int {
	l, r := 0, len(arr)-k
	for l < r {
		mid := (l + r) / 2
		if x-arr[mid] > arr[mid+k]-x {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return arr[l : l+k]
}

func findClosestElementsMy(arr []int, k int, x int) []int {
	index := 0
	index = helper(arr, x)
	fmt.Println(index, arr[index], x)
	res := make([]int, 0, k)
	for l, r := index, index+1; k > 0; k-- {
		if l > -1 && r < len(arr) {
			if x-arr[l] <= arr[r]-x {
				hash := make([]int, len(res)+1)
				hash[0] = arr[l]
				copy(hash[1:], res)
				res = hash
				l--
			} else {
				res = append(res, arr[r])
				r++
			}
		} else if l > -1 {
			hash := make([]int, len(res)+1)
			hash[0] = arr[l]
			copy(hash[1:], res)
			res = hash
			l--
		} else if r < len(arr) {
			res = append(res, arr[r])
			r++
		}
	}
	return res
}

func helper(arr []int, x int) int {
	il := bsLeft(0, len(arr)-1, arr, x)
	ir := bsRight(0, len(arr)-1, arr, x)
	if x-arr[ir] < arr[il]-x {
		return ir
	}
	return il
}

func bsLeft(l, r int, arr []int, x int) int {
	for l < r {
		m := (l + r) / 2
		if arr[m] > x {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}

func bsRight(l, r int, arr []int, x int) int {
	for l < r {
		m := (l + r + 1) / 2
		if arr[m] <= x {
			l = m
		} else {
			r = m - 1
		}
	}
	return l
}
