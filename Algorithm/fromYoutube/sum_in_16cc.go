package main

import "fmt"

func main() {
	n1, n2 := "19", "100"
	fmt.Println(sumIn16cc(n1, n2))
}

func sumIn16cc(n1 string, n2 string) string {
	dataMap := map[string]int{"0": 0,
		"1": 1, "2": 2, "3": 3, "4": 4, "5": 5,
		"6": 6, "7": 7, "8": 8, "9": 9, "a": 10,
		"b": 11, "c": 12, "d": 13, "e": 14, "f": 15}
	dataSlice := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f"}

	reverse := ""
	n, m := len(n1), len(n2)
	if m > n {
		m, n = n, m
		n1, n2 = n2, n1
	}
	buf := "0"
	for i := 0; i < m; i++ {
		sum := dataMap[string(n1[n-1-i])] + dataMap[string(n2[m-1-i])] + dataMap[buf]
		reverse += dataSlice[sum%16]
		buf = dataSlice[sum/16]
	}
	for i := m; i < n; i++ {
		sum := dataMap[string(n1[n-1-i])] + dataMap[buf]
		reverse += dataSlice[sum%16]
		buf = dataSlice[sum/16]
	}
	if buf != "0" {
		reverse += buf
	}

	res := ""
	for i := len(reverse) - 1; i > -1; i-- {
		res += string(reverse[i])
	}

	return res
}
