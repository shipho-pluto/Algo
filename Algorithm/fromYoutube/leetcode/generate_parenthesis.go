package main

import "fmt"

func main() {
	n := 3
	fmt.Println(generateParenthesis(n))
}

func generateParenthesis(n int) []string {
	var res []string
	rec(&res, "", 0, 0, n)
	return res
}

func rec(res *[]string, s string, l, r, n int) {
	if len(s) == 2*n {
		*res = append(*res, s)
		return
	}

	if l < n {
		rec(res, s+"(", l+1, r, n)
	}
	if l > r {
		rec(res, s+")", l, r+1, n)
	}

}
