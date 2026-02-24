package main

import "fmt"

func main() {
	s1 := "ab"
	s2 := "eidboaooo"
	fmt.Println(onePerm(s1, s2))
}

func onePerm(s1 string, s2 string) bool {
	n := len(s1)
	maskEx := [26]int{}
	maskTemp := [26]int{}
	for i := range s1 {
		maskEx[s1[i]-'a']++
		maskTemp[s2[i]-'a']++
	}

	eq := 0
	for i := range 26 {
		if maskEx[i] == maskTemp[i] {
			eq++
		}
	}
	if eq == 26 {
		return true
	}
	for i := n; i < len(s2); i++ {
		maskTemp[s2[i-n]-'a']--
		if maskEx[s2[i-n]-'a'] == maskTemp[s2[i-n]-'a'] {
			eq++
		} else {
			eq--
		}

		maskTemp[s2[i]-'a']++
		if maskEx[s2[i]-'a'] == maskTemp[s2[i]-'a'] {
			eq++
		} else {
			eq--
		}

		if eq == 26 {
			return true
		}

	}
	return false
}
