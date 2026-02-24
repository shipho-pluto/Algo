package main

import "fmt"

func main() {
	arr := []int{1, -1, 3, 2, 3, -5, 0, -3}
	k := 0
	fmt.Println(cntSumEqK(arr, k))

	arr = []int{2, 2, 1, 1, 0, 2, -1}
	k = 2
	fmt.Println(cntSumEqK(arr, k))

	arr = []int{2, 2, 1, 1, 0, 2}
	k = 2
	fmt.Println(cntSumEqKLR(arr, k))
}

func cntSumEqK(arr []int, k int) int {
	cnt := 0
	st := make(map[int]int)
	st[0] = 1
	sum := 0
	for i := range arr {
		sum += arr[i]
		if v, ok := st[sum-k]; ok {
			cnt += v
		}
		st[sum]++
	}

	return cnt
}

func cntSumEqKLR(arr []int, k int) int {
	cnt := 0
	sum := 0
	for l, r := 0, 0; l < len(arr); {
		if r < len(arr) {
			if sum+arr[r] < k {
				sum += arr[r]
				r++
			} else if sum+arr[r] == k {
				cnt++
				sum += arr[r]
				r++
			} else {
				sum -= arr[l]
				l++
				if sum == k {
					cnt++
				}
			}
		} else {
			if sum-arr[l] < k {
				break
			} else if sum-arr[l] == k {
				cnt++
			}
			sum -= arr[l]
			l++
		}
	}

	return cnt
}
