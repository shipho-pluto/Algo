package main

import "fmt"

func main() {
	arr := []int{1, 3, 5, 7, 2, 3, 6, 10, 8, 13}
	x := 13
	m := make(map[int]int)
	for i, v := range arr {
		if num, ok := m[x-v]; ok {
			fmt.Println(arr[num], v)
			break
		}
		m[v] = i
	}

	dictionary := []string{
		"aboba", "roar", "bound", "fade", "light", "power", "graduation", "lop", "donda",
	}
	M := make(map[string]int)
	for _, v := range dictionary {
		M[v]++
		for i := range v {
			s := v[:i] + v[i+1:]
			M[s]++
		}
	}
	fmt.Println(M)
	text := []string{
		"aboba", "abob", "abo", "dona", "fad", "fadee", "roar", "roa",
	}
	cnt := 0
	for _, v := range text {
		if _, ok := M[v]; ok {
			cnt++
		}
	}
	fmt.Println(cnt)
}
