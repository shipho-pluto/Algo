package main

import "fmt"

func Histogram(s string) {
	m := make(map[rune]int)
	n := len(s)
	maxV := 0
	for i := range n {
		m[rune(s[i])]++
		maxV = max(maxV, m[rune(s[i])])
	}
	res := make([]string, maxV+1)
	for k, v := range m {
		for i := 0; i < maxV; i++ {
			if i >= maxV-v {
				res[i] += "#"
			} else {
				res[i] += " "
			}
		}
		res[maxV] += string(k)
	}
	for _, v := range res {
		fmt.Println(v)
	}
}

func main() {
	s := "Hello world!"
	Histogram(s)
}
