package main

import "fmt"

func main() {
	arr := []int{1, 0, 1, 2, 0, 3, 4, 0, 5, 6, 7, 0, 8}
	fmt.Println(moveZeroes(arr))

	arr = []int{1, -1, 4, 5, 10, 23, 11, -3}
	k := 8
	fmt.Println(subSumK(arr, k))

	arr = []int{1, 4, 5, 0, 10, 23, 11}
	k = 15
	fmt.Println(subSumKLR(arr, k))

	word := "abcca"
	fmt.Println(palin2(word))

}

func palin2(word string) bool {
	for l, r := 0, len(word)-1; l < r; l, r = l+1, r-1 {
		if word[l] != word[r] {
			if word[l+1] != word[r] && word[l] != word[r-1] {
				return false
			} else if word[l+1] == word[r] {
				l++
			} else {
				r--
			}
		}
	}
	return true
}

func subSumKLR(arr []int, k int) (int, int) {
	sum := arr[0]
	for l, r := 0, 1; r < len(arr); {
		if sum+arr[r] < k {
			sum += arr[r]
			r++
		} else if sum+arr[r] > k {
			sum -= arr[l]
			l++
		} else {
			return l, r
		}
	}
	return -1, -1
}

func subSumK(arr []int, k int) (int, int) {
	st := make(map[int]int)
	sum := 0
	for i := range arr {
		sum += arr[i]
		if l, ok := st[sum-k]; ok {
			return l, i
		}

		st[arr[i]] = i + 1
	}
	return -1, -1
}

func moveZeroes(arr []int) []int {
	for l := 0; l < len(arr); l++ {
		r := l + 1
		for ; r < len(arr) && arr[r] == 0; r++ {
		}
		if arr[l] == 0 && r < len(arr) {
			arr[l], arr[r] = arr[r], arr[l]
		}
	}

	return arr
}
