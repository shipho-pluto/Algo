package main

import "fmt"

func theSame(s1, s2 string) bool {
	n, m := len(s1), len(s2)
	if n-m > 1 || m-n > 1 {
		return false
	}
	if n > m {
		n, m = m, n
		s1, s2 = s2, s1
	}
	eq := 0
	mask1, mask2 := [26]int{}, [26]int{}
	for i := 0; i < n; i++ {
		mask1[s1[i]-'a']++
		mask2[s2[i]-'a']++
	}
	for i := n; i < m; i++ {
		mask2[s2[i]-'a']++
	}

	for i := range 26 {
		if mask1[i] == mask2[i] {
			eq++
		}
	}
	if n == m {
		return eq == 26 || eq == 24
	}
	return eq == 25
}

func main() {
	s1, s2 := "cat", "cat"
	fmt.Println(theSame(s1, s2))
}
