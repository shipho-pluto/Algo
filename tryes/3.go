package main

import "fmt"

func main() {
	a := "abs"
	b := "abrr"
	fmt.Println(oneEdit(a, b))

	arr := []int{0, 0, 1, 1, 0, 1, 1, 0, 1, 0, 0, 1, 1, 1}
	k := 3
	fmt.Println(maxLenOf1WithK0(arr, k))

	nums := []int{1, 0, 3, 4, 3, 2, 5, 9, 10, 1, 0}
	l, r := []int{0, 4, 6}, []int{2, 7, 8}
	fmt.Println(changesAdd1(nums, l, r))
	fmt.Println(changesAdd1New(nums, l, r))
}

func changesAdd1New(arr []int, l []int, r []int) any {
	c := 0
	for li, ri, i := 0, 0, 0; i < len(arr); i++ {
		if li < len(l) && l[li] == i {
			c++
		}
		if ri < len(r) && r[ri] == i {
			c--
		}
		arr[i] += c
	}

	return arr
}

func changesAdd1(nums []int, l []int, r []int) any {
	pref := make([]int, len(nums)+1)
	for li, ri, i := 0, 0, 0; i < len(nums); i++ {
		if li < len(l) && i == l[li] {
			pref[i]++
			li++
		} else if ri < len(r) && i == r[ri] {
			pref[i]--
			ri++
		}

		pref[i+1] = pref[i]
		nums[i] += pref[i]
	}

	return nums
}

func q(x bool) int {
	if x {
		return 1
	}
	return 0
}

func maxLenOf1WithK0(arr []int, k int) int {
	kc := k + 1
	r := 0
	for r = range arr {
		if arr[r] == 0 {
			if kc-1 == 0 {
				break
			}
			kc--
		}
	}
	res := r
	for l := 0; r < len(arr); r++ {
		if arr[r] == 0 {
			kc--
			for ; l < r; l++ {
				if kc > 0 {
					break
				}
				kc += q(arr[l] == 0)
			}
		} else {
			res = max(r-l+1, res)
		}

	}
	return res
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

func oneEdit(a string, b string) bool {
	n, m := len(a), len(b)
	if abs(n-m) > 1 || n == 0 && m == 0 {
		return false
	}

	if m < n {
		n, m = m, n
		a, b = b, a
	}

	eq := 0
	if n == m {
		for i := range n {
			if a[i] != b[i] {
				eq++
			}
			if eq == 2 {
				return false
			}
		}
		return true
	} else {
		for l, r := 0, 0; l < n; l, r = l+1, r+1 {
			if a[l] != b[r] {
				r++
				if a[l] != b[r] {
					return false
				}
			}
			eq++
		}
		return eq == n
	}
}
