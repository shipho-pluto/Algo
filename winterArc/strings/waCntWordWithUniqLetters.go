package main

import "fmt"

func main() {
	s := "abcabcabc"
	fmt.Println(cntUniq(s))
}

func cntUniq(s string) int {

	i := 0
	mask := [26]bool{}
	for ; i < len(s); i++ {
		if mask[s[i]-'a'] {
			break
		}
		mask[s[i]-'a'] = true
	}

	cnt := i * (i + 1) / 2
	l := 0
	for r := i; r < len(s); r++ {
		if mask[s[r]-'a'] {
			for ; l < r; l++ {
				mask[s[l]-'a'] = false
				if !mask[s[r]-'a'] {
					break
				}
			}
		} else {
			cnt += r - l
		}
		mask[s[r]-'a'] = true
	}

	return cnt
}
