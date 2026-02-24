package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	arr := []int{1, 3, 5, 8}
	fmt.Println(area(arr))

	arr = []int{4, 3, 2, 1, 5, 4, 3, 1, 1, 3}
	fmt.Println(high(arr))

	a, b := []int{1, 4, 5, 6, 0, 0, 0}, []int{-1, 3, 7}
	fmt.Println(mergeInA(a, b, 4, 3))

	s := "aaabbdbdggff"
	fmt.Println(rle(s))

}

func rle(s string) any {
	ns := strings.Builder{}
	l, r := 0, 1
	for ; r < len(s); r++ {
		if s[r] != s[r-1] {
			if r-l > 1 {
				ns.WriteString(strconv.Itoa(r-l) + string(s[r-1]))
			} else {
				ns.WriteString(string(s[r-1]))
			}
			l = r
		}
	}
	if s[r-1] == s[r-2] {
		ns.WriteString(strconv.Itoa(r-l) + string(s[r-1]))
	} else {
		ns.WriteString(string(s[r-1]))
	}

	return ns.String()
}

func mergeInA(a []int, b []int, k int, n int) any {
	for l, r, i := k-1, n-1, len(a)-1; i > -1; i-- {
		if l > -1 && (r > -1 && a[l] > b[r] || r == -1) {
			a[i] = a[l]
			l--
		} else {
			a[i] = b[r]
			r--
		}
	}
	return a
}

func high(arr []int) any {
	maxI := 0
	maxE := arr[0]
	for i := range arr {
		if arr[i] > maxE {
			maxE = arr[i]
			maxI = i
		}
	}
	res := 0
	mb, ma := arr[0], arr[len(arr)-1]
	for l, r := 0, len(arr)-1; l < maxI || r > maxI; l, r = l+1, r-1 {
		if l < maxI {
			mb = max(mb, arr[l])
			if arr[l] < mb {
				res += mb - arr[l]
			}
		}

		if r > maxI {
			ma = max(ma, arr[r])
			if arr[r] < ma {
				res += ma - arr[r]
			}
		}
	}

	return res
}

func area(arr []int) any {
	res := 0
	for l, r := 0, len(arr)-1; l < r; {
		res = max(res, (r-l)*min(arr[r], arr[l]))
		if arr[l] < arr[r] {
			l++
		} else {
			r--
		}
	}

	return res
}
