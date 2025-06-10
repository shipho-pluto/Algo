package main

import (
	"fmt"
)

func function(a, b string) string {
	st := map[rune]int{
		'0': 0, '1': 1, '2': 2, '3': 3,
		'4': 4, '5': 5, '6': 6, '7': 7,
		'8': 8, '9': 9, 'a': 10, 'b': 11,
		'c': 12, 'd': 13, 'e': 14, 'f': 15}
	con := []rune{
		'0', '1', '2', '3',
		'4', '5', '6', '7',
		'8', '9', 'a', 'b',
		'c', 'd', 'e', 'f'}
	ac := []rune(a)
	bc := []rune(b)

	rest := 0
	ra := len(ac) - 1
	rb := len(bc) - 1
	s := ""
	for ra >= 0 && rb >= 0 {
		s1 := st[ac[ra]]
		s2 := st[bc[rb]]
		sum := s1 + s2 + rest
		rest = sum / 16
		s += string(con[sum%16])
		ra--
		rb--
	}
	for ra >= 0 {
		s1 := st[ac[ra]]
		sum := s1 + rest
		rest = sum / 16
		s += string(con[sum%16])
		ra--
	}
	for rb >= 0 {
		s2 := st[bc[rb]]
		sum := s2 + rest
		rest = sum / 16
		s += string(con[sum%16])
		rb--
	}
	if rest != 0 {
		s += string(con[rest])
	}
	newS := ""
	for i := len([]rune(s)) - 1; i > -1; i-- {
		newS += string(s[i])
	}
	return newS
}

func main() {
	var s1, s2 string

	fmt.Scan(&s1)
	fmt.Scan(&s2)
	fmt.Println(function(s1, s2))
}
