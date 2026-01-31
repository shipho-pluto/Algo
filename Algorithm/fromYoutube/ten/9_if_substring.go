package main

import "fmt"

func main() {
	s, f := "cat", "car"
	fmt.Println(ifSubstring(s, f))
	s, f = "cat", "ca"
	fmt.Println(ifSubstring(s, f))
	s, f = "cat", "catt"
	fmt.Println(ifSubstring(s, f))
	s, f = "cat", "cat"
	fmt.Println(ifSubstring(s, f))
}

func ifSubstring(s, f string) bool {
	n, m := len(s), len(f)
	if m > n {
		m, n = n, m
		s, f = f, s
	}
	if n-m > 1 {
		return false
	}
	eq := 0
	mask, wind := [26]int{}, [26]int{}
	for i := range m {
		wind[s[i]-'a']++
		mask[f[i]-'a']++
	}
	for i := m; i < n; i++ {
		wind[s[i]-'a']++
	}
	for i := range 26 {
		if mask[i] == wind[i] {
			eq++
		}
	}
	if n == m {
		return eq == 26 || eq == 24
	}
	return eq == 25
}
