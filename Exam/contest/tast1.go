package main

import "fmt"

func main() {
	var s string
	fmt.Scanln(&s)
	isR := false
	for i := range s {
		if s[i] == 'R' {
			isR = true
		} else if s[i] == 'M' {
			break
		}
	}

	if isR {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
