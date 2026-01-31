package main

import "fmt"

func OneEditApart(a, b string) bool {
	n, m := len(a), len(b)
	if n-m > 1 || m-n > 1 {
		return false
	}
	alA, alB := [26]int{}, [26]int{}
	if n == m {
		eq := 0
		for i := 0; i < n; i++ {
			if a[i] != b[i] {
				eq++
			}
		}
		return 0 < eq && eq < 2
	}
	if n > m {
		n, m = m, n
		b, a = a, b
	}
	for i := 0; i < m; i++ {
		alB[int(b[i]-'a')]++
		if i < n {
			alA[int(a[i]-'a')]++
		}
	}
	eq := 0
	for i := 0; i < 26; i++ {
		if alA[i] == alB[i] {
			eq++
		}
	}
	return eq == 25
}

func main() {
	fmt.Println(OneEditApart("cap", "cat"))
}
