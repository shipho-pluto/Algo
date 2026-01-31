package main

import "fmt"

func main() {
	arr := []int{0, 1, 2, 3, 4, 10, 0, 2, 1, -1, -2, -10, -11, -12}
	fmt.Println(maxLenMonotonicArray(arr))
}

func maxLenMonotonicArray(arr []int) (int, int) {
	res := []int{0, 0, 0}
	for monotonic, cnt, i := true, 1, 1; i < len(arr); {
		if monotonic && arr[i] < arr[i-1] {
			cnt++
			i++
		} else if monotonic && arr[i] > arr[i-1] {
			cnt = 1
			monotonic = false
		} else if !monotonic && arr[i] > arr[i-1] {
			cnt++
			i++
		} else if !monotonic && arr[i] < arr[i-1] {
			cnt = 1
			monotonic = true
		} else if arr[i] == arr[i-1] {
			cnt = 1
			i++
		}

		if res[0] < cnt {
			res = []int{cnt, i - cnt, i - 1}
		}
	}

	return res[1], res[2]
}
