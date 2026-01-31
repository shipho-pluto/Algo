package main

import "fmt"

func main() {
	dictionary := []string{
		"aboba", "roar", "bound", "fade", "light", "power", "graduation", "lop", "donda",
	}
	text := []string{
		"aboba", "abob", "abo", "dona", "fad", "fadee", "roar", "roa",
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
