package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n, m := len(nums1), len(nums2)
	l, r := 0, 0
	res := make([]int, n+m)
	sum := 0
	for i := 0; l < n || r < m; i++ {
		if l < n && r < m {
			if nums1[l] < nums2[r] {
				res[i] = nums1[l]
				l++
			} else {
				res[i] = nums2[r]
				r++
			}
		} else if l < n {
			res[i] = nums1[l]
			l++
		} else {
			res[i] = nums2[r]
			r++
		}
		sum += res[i]

	}
	if (n+m)%2 == 1 {
		return float64(res[(n+m)/2])
	}
	return float64(res[(n+m)/2]+res[(n+m)/2-1]) / 2
}

func main() {
	arr1, arr2 := []int{1, 2}, []int{3, 4}
	fmt.Println(findMedianSortedArrays(arr1, arr2))
}
