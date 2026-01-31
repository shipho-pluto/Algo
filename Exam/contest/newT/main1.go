package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var n int
	fmt.Scan(&n)
	mask := [10]int{}

	for n != 0 {
		mask[n%10]++
		n /= 10
	}

	res := strings.Builder{}

	if mask[0] == 0 {
		for i := range mask {
			for range mask[i] {
				res.WriteString(strconv.Itoa(i))
			}
		}
	} else {
		for i := range mask {
			if mask[i] != 0 && i != 0 {
				res.WriteString(strconv.Itoa(i))
				mask[i]--
				break
			}
		}
		for i := range mask {
			for range mask[i] {
				res.WriteString(strconv.Itoa(i))
			}
		}
	}

	fmt.Println(res.String())
}
