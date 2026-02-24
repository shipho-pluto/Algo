package main

import "fmt"

func main() {
	dictionary := []string{
		"i", "am", "have", "you", "we", "are", "siblings", "program", "language", "go", "work",
	}
	text := []string{
		"i", "am", "yuo", "programm", "languje", "go", "works",
	}

	fmt.Println(cntRightWords(dictionary, text))
}

func cntRightWords(d []string, t []string) int {
	cnt := 0

	stForDict := make(map[[26]int]string)
	for i := range d {
		maskForWord := [26]int{}
		for j := range d[i] {
			maskForWord[d[i][j]-'a']++
		}
		stForDict[maskForWord] = d[i]
	}

	for i := range t {
		maskForWord := [26]int{}
		for j := range t[i] {
			maskForWord[t[i][j]-'a']++
		}
		if word, ok := stForDict[maskForWord]; ok && t[i] == word {
			cnt++
		} else {
			for k := range 26 {
				maskForWord[k]++
				if _, ok = stForDict[maskForWord]; ok {
					cnt++
					break
				}
				maskForWord[k] -= 2
				if _, ok = stForDict[maskForWord]; ok {
					cnt++
					break
				}
				maskForWord[k]++
			}
		}
	}

	return cnt
}
