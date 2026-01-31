package main

import "fmt"

func returnIf(a, f int) int {
	if a == int(int32(a)) {
		return f * a
	}
	if f == -1 {
		return -2147483648
	}
	return 2147483647
}

func myAtoi(s string) int {
	n := len(s)
	integer := 0
	flag := 1
	minusTo := 0
	spaceTo := 0
	for i := 0; i < n; i++ {
		if integer >= 2147483648 {
			if flag == -1 {
				return -2147483648
			}
			return 2147483647
		}
		if s[i] == ' ' {
			if spaceTo == 1 {
				return returnIf(integer, flag)
			}
			minusTo++
		} else if s[i] == '+' {
			if i == minusTo {
				spaceTo = 1
			} else {
				return returnIf(integer, flag)
			}
		} else if s[i] == '-' {
			if i == minusTo {
				flag = -1
			} else {
				return returnIf(integer, flag)
			}
			spaceTo = 1
		} else if s[i] == '0' || s[i] == '1' || s[i] == '2' || s[i] == '3' || s[i] == '4' ||
			s[i] == '5' || s[i] == '6' || s[i] == '7' || s[i] == '8' || s[i] == '9' {
			integer *= 10
			integer += int(s[i] - '0')
			spaceTo = 1
		} else {
			return returnIf(integer, flag)
		}
	}

	return returnIf(integer, flag)
}

func main() {
	s := "18446744073709551617"
	fmt.Println(myAtoi(s))
}
