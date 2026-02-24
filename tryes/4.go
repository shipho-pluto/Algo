package main

import "fmt"

func main() {
	a, b := "ab", "arrtghrfdba"
	fmt.Println(permutationInString(a, b))

	s := "aaaaaaasddd"
	fmt.Println(longestSub(s))
	fmt.Println(longestSubLR(s, 3))
}

type Tree struct {
	val   int
	left  *Tree
	right *Tree
}

func walk(root *Tree, loop string, t1, t2 int, arr [2]string) {
	if root != nil {
		if root.val == t1 && arr[0] == "" {
			arr[0] = loop
		} else if root.val == t2 && arr[1] == "" {
			arr[1] = loop
		}
		if arr[0] != "" && arr[1] != "" {
			return
		}
		walk(root.left, loop+"0", t1, t2, arr)
		walk(root.right, loop+"1", t1, t2, arr)
	}
}

func longestSub(s string) any {
	st := make(map[uint8]int)
	res := 0
	p := 0
	for r := range s {
		p = max(p, st[s[r]-'a'])
		res = max(res, r-p)
		st[s[r]-'a'] = r
	}

	return res
}

func longestSubLR(s string, k int) any {
	st := [26]int{}
	res := 0
	for l, r := 0, 0; r < len(s); r++ {
		if st[s[r]-'a'] == k {
			for ; l < r && st[s[r]-'a'] == k; l++ {
				st[s[l]-'a']--
			}
		}
		res = max(res, r-l+1)
		st[s[r]-'a']++
	}
	return res
}

func permutationInString(a string, b string) any {
	mask, wind := [26]int{}, [26]int{}
	for i := range a {
		mask[a[i]-'a']++
		wind[b[i]-'a']++
	}
	eq := 0
	for i := range 26 {
		if mask[i] == wind[i] {
			eq++
		}
	}
	if eq == 26 {
		return true
	}
	for i := len(a); i < len(b); i++ {
		if mask[b[i-len(a)]-'a'] == wind[b[i-len(a)]-'a'] {
			eq--
		}
		wind[b[i-len(a)]-'a']--
		if mask[b[i-len(a)]-'a'] == wind[b[i-len(a)]-'a'] {
			eq++
		}

		if mask[b[i]-'a'] == wind[b[i]-'a'] {
			eq--
		}
		wind[b[i]-'a']++
		if mask[b[i]-'a'] == wind[b[i]-'a'] {
			eq++
		}

		if eq == 26 {
			return true
		}
	}

	return false
}
