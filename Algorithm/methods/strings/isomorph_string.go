package main

import "fmt"

func main() {
	s := "tatacc"
	t := "papacc"
	fmt.Println(isomorphString(s, t))
}

func isomorphString(s string, t string) bool {
	n, m := len(s), len(t)
	if n != m {
		return false
	}

	mask1, mask2 := make(map[uint8]int), make(map[uint8]int)
	for i := 0; i < n; i++ {
		if _, ok := mask1[s[i]-'a']; ok && mask1[s[i]-'a'] != int(t[i]-'a') {
			return false
		}
		if _, ok := mask1[t[i]-'a']; ok && mask1[t[i]-'a'] != int(s[i]-'a') {
			return false
		}
		mask1[s[i]-'a'] = int(t[i] - 'a')
		mask2[t[i]-'a'] = int(s[i] - 'a')
	}
	return true
}
