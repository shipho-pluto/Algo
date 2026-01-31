package main

import (
	"fmt"
)

func main() {
	arr := []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}
	fmt.Println(compress(arr))
	for _, i := range arr {
		fmt.Println(string(i), i)
	}
}
func compress(arr []byte) int {
	cnt := 1
	n := len(arr)
	if n < 2 {
		return cnt
	}
	res := make([]byte, 0, n)

	for l, r := 0, 1; r < n; r++ {
		if arr[r] == arr[r-1] {
			cnt++
		} else {
			if cnt == 1 {
				res = append(res, arr[l])
			} else if cnt < 10 {
				res = append(res, arr[l], byte('A'-17+cnt))
			} else {
				res = append(res, arr[l])
				var rev []byte
				for cnt != 0 {
					rev = append(rev, byte('A'-17+cnt%10))
					cnt /= 10
				}
				for i := len(rev) - 1; i >= 0; i-- {
					res = append(res, rev[i])
				}
			}
			l = r
			cnt = 1
		}
	}

	if arr[n-1] == arr[n-2] {
		if cnt < 10 {
			res = append(res, arr[n-1], byte('A'-17+cnt))
		} else {
			res = append(res, arr[n-1])
			var rev []byte
			for cnt != 0 {
				rev = append(rev, byte('A'-17+cnt%10))
				cnt /= 10
			}
			for i := len(rev) - 1; i >= 0; i-- {
				res = append(res, rev[i])
			}
		}
	} else {
		res = append(res, arr[n-1])
	}

	copy(arr, res)

	return len(res)
}
