package main

import "fmt"

func Rock(a, b string) int {
	m, n := len(a), len(b)
	alA := [26]int{}
	alB := [26]int{}
	for i := 0; i < m; i++ {
		alA[a[i]-'a']++
	}
	for i := 0; i < n; i++ {
		alB[b[i]-'a']++
	}
	sum := 0
	for i := 0; i < 26; i++ {
		if alA[i] != 0 {
			sum += alB[i]
		}
	}

	return sum
}

func main() {
	var a, b string
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Println(Rock(a, b))
}
