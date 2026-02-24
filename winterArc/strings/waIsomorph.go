package main

import "fmt"

func main() {
	s := "tatacc"
	t := "totacc"
	fmt.Println(isomorph(s, t))
}

func isomorph(s string, t string) bool {
	mask := [26]uint8{}
	if len(s) != len(t) {
		return false
	}
	for i := range s {
		if mask[s[i]-'a'] != 0 {
			if s[i]-t[i]+126 != mask[s[i]-'a'] {
				return false
			}
		}
		mask[s[i]-'a'] = s[i] - t[i] + 126
	}
	return true
}
