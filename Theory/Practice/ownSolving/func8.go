package main

import "fmt"

func abs8(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func OneEditApart(s1, s2 string) bool {
	n, m := len(s1), len(s2)
	if abs8(n-m) > 1 {
		return false
	}

	m1, m2 := make([]int, 26), make([]int, 26)
	eq := 0
	if n == m {
		for i := 0; i < n; i++ {
			m1[s1[i]-'a']++
			m2[s2[i]-'a']++
		}
		for i := range 26 {
			if m1[i] != m2[i] {
				eq++
			}
			if eq > 2 {
				return false
			}
		}
		return eq == 0 || eq == 2
	}
	if m > n {
		m, n = n, m
		s1, s2 = s2, s1
	}
	for i := range m {
		m1[s1[i]-'a']++
		m2[s2[i]-'a']++
	}
	for i := m; i < n; i++ {
		m1[s1[i]-'a']++
	}
	for i := range 26 {
		if m1[i] != m2[i] {
			eq++
		}
		if eq > 1 {
			return false
		}
	}
	return eq == 1

}

func main() {
	s1, s2 := "car", "cde"
	fmt.Println(OneEditApart(s1, s2))
}
