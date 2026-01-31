package main

import "fmt"

func is(q bool) int {
	if q {
		return 1
	}
	return 0
}

func StringInString(s1, s2 string) int {
	n, m := len(s1), len(s2)
	if n < m {
		return -1
	}
	al1, al2 := make([]int, 26), make([]int, 26)
	eq := 0
	for i := range s2 {
		al2[s2[i]-'a']++
		al1[s1[i]-'a']++
	}
	for i := range 26 {
		if al2[i] == al1[i] {
			eq++
		}
	}
	if eq == 26 {
		return 0
	}
	for i := m; i < n; i++ {
		if al1[s1[i-m]-'a'] == al2[s1[i-m]-'a'] {
			eq--
		}
		al1[s1[i-m]-'a']--
		if al1[s1[i-m]-'a'] == al2[s1[i-m]-'a'] {
			eq++
		}

		if al1[s1[i]-'a'] == al2[s1[i]-'a'] {
			eq--
		}
		al1[s1[i]-'a']++
		if al1[s1[i-m]-'a'] == al2[s1[i-m]-'a'] {
			eq++
		}
		if eq == 26 {
			return i
		}
	}

	return -1
}

func main() {
	s1, s2 := "abbbabababaadasasdafafas", "sad"
	fmt.Println(StringInString(s1, s2))
}
