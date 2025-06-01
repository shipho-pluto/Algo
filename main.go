package main

import (
	"fmt"
	"os"
)

func main() {
	var k1, m, k2, p2, n2 int
	fmt.Fscanln(os.Stdin, &k1, &m, &k2, &p2, &n2)
	if (p2-1)*m >= k2 && p2 > 1 || n2 > k2 && p2 == 1 {
		fmt.Println(-1, -1)
		return
	}
	kVm := k2 / n2
	if kVm*n2 != k2 {
		kVm++
	}
	kVp := kVm * m
	p1 := k1 / kVp
	if p1*kVp != k1 {
		p1++
	}
	n1 := (k1 - kVp*(p1-1)) / kVm
	if n1*kVm != (k1 - kVp*(p1-1)) {
		n1++
	}
	fmt.Println(p1, n1)
}
