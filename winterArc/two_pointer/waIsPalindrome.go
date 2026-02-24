package main

import "fmt"

func main() {
	p := "abooba"
	fmt.Println(isPalindrome(p))
}

func isPalindrome(p string) bool {
	n := len(p)
	for l, r := 0, n-1; l < r; l, r = l+1, r-1 {
		if p[l] != p[r] {
			return false
		}
	}
	return true
}
