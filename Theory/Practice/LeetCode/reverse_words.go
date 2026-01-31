package main

import (
	"fmt"
)

func reverseWords(s string) string {
	var words []string
	index := 0
	for i := range len(s) {
		if s[i] == ' ' {
			words = append(words, s[index:i])
			index = i + 1
		}
	}
	words = append(words, s[index:])

	n := len(words)
	for i := range n {
		newWord := ""
		for j := len(words[i]) - 1; j > -1; j-- {
			newWord += string(words[i][j])
		}
		words[i] = newWord
	}
	res := ""
	for _, v := range words {
		res += v + " "
	}
	return res
}

func main() {
	s := "mr ding dong in my bell"
	fmt.Println(reverseWords(s))
}
