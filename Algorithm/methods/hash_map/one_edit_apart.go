package main

import "fmt"

func OneEditApart(s1, s2 string) bool {
	n, m := len(s1), len(s2)

	if m > n {
		n, m = m, n
		s1, s2 = s2, s1
	}

	if n-m > 1 {
		return false
	}

	eq := 0
	mask1, mask2 := [26]int{}, [26]int{}
	for i := 0; i < m; i++ {
		mask1[s1[i]-'a']++
		mask2[s2[i]-'a']++
	}
	for i := m; i < n; i++ {
		mask1[s1[i]-'a']++
	}

	for i := range 26 {
		if mask1[i] == mask2[i] {
			eq++
		}
	}
	if n == m {
		return eq == 24 || eq == 26
	}
	return eq == 25
}

func main() {
	s1, s2 := "cat", "ttac"
	fmt.Println(OneEditApart(s1, s2))
}
