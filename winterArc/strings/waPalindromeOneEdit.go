package main

import "fmt"

func main() {
	a := "abccba"
	fmt.Println(palindrome(a))

}

func palindrome(a string) bool {
	for l, r := 0, len(a)-1; l < r; l, r = l+1, r-1 {
		if a[l] != a[r] {
			if a[l+1] != a[r] && a[l] != a[r-1] {
				return false
			} else if a[l+1] == a[r] {
				l++
			} else {
				r--
			}
		}
	}

	return true
}
