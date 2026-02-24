package main

import "fmt"

func main() {
	ar := "000XXOOYOOX0"
	fmt.Println(nearElementToTarget(ar))
}

func nearElementToTarget(arr string) int {
	l := 0
	isX := false
	for ; l < len(arr); l++ {
		if arr[l] == 'X' {
			isX = true
			break
		} else if arr[l] == 'Y' {
			isX = false
			break
		}
	}

	minDif := len(arr)
	for r := l + 1; r < len(arr); r++ {
		if isX && arr[r] == 'Y' {
			minDif = min(minDif, r-l)
			l = r
			isX = false
		} else if isX && arr[r] == 'X' {
			l = r
		}

		if !isX && arr[r] == 'X' {
			minDif = min(minDif, r-l)
			l = r
			isX = true
		} else if !isX && arr[r] == 'Y' {
			l = r
		}
	}

	return minDif
}
