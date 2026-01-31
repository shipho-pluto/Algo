package main

import "fmt"

func sum16ccNum(n1, n2 string) string {
	n, m := len(n1), len(n2)
	res := ""
	if m > n {
		m, n = n, m
		n1, n2 = n2, n1
	}

	data := map[string]int{
		"a": 10, "b": 11, "c": 12, "d": 13, "e": 14, "f": 15,
		"0": 0, "1": 1, "2": 2, "3": 3, "4": 4,
		"5": 5, "6": 6, "7": 7, "8": 8, "9": 9,
	}
	convert := []string{
		"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f",
	}
	val := "0"
	for i := 0; i < m; i++ {
		sum := data[string(n1[n-1-i])] + data[string(n2[m-1-i])] + data[val]
		res += convert[sum%16]
		val = convert[sum/16]
	}
	for i := m; i < n; i++ {
		sum := data[string(n1[n-1-i])] + data[val]
		res += convert[sum%16]
		val = convert[sum/16]
	}
	if val != "0" {
		res += val
	}
	result := ""
	for i := len(res) - 1; i > -1; i-- {
		result += string(res[i])
	}
	return result
}

func main() {
	s1, s2 := "19", "1"
	fmt.Println(sum16ccNum(s1, s2))
}
