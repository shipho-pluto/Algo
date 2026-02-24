package main

import "fmt"

func main() {
	dictionary := []string{
		"i", "am", "have", "you", "we", "are", "siblings", "program", "language", "go", "work",
	}
	text := []string{
		"i", "am", "yuo", "programm", "languje", "go", "works",
	}
	fmt.Println(cntRightWords(text, dictionary))
}

func cntRightWords(text []string, dictionary []string) int {
	pref := make(map[string]bool)
	for i := 0; i < len(dictionary); i++ {
		pref[dictionary[i]] = true
		for j := 0; j < len(dictionary[i]); j++ {
			pref[dictionary[i][:j]+dictionary[i][j+1:]] = true
		}
	}
	res := 0
	for _, word := range text {
		if pref[word] {
			res++
		}
	}

	return res
}
