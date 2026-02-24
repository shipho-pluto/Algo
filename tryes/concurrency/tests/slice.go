package main

import "fmt"

func main() {
	a := []int{2, 3, 5, 7, 9} // a = [2, 3, 5, 7, 9]
	println(cap(a))           // l=5, c=5
	b := a[1:4]               // b = [3, 5, 7, nil]
	println(cap(b))           // l=3, c=4

	A := make([]int, 3, 5)                      // [0, 0, 0, nil, nil]
	fmt.Printf("%v %d %d\n", A, len(A), cap(A)) // [0, 0, 0], 3, 5

	a = []int{}
	println(cap(a)) // 0

	a = []int{2, 3, 5, 7, 9}
	b = a[1:4] // b = [3, 5, 7] l = 3, c = 4
	// println(b[3]) // panic
	println(b[:4][3]) // b[:4] = [3, 5, 7, "9"] b[:4][3] = 9

	primes := []int{2, 3, 5, 7, 11, 13} // l=6, c=6
	println(sum(primes[0:3]))           // [2, 3, 5] l=3, c=6 -> primes[a] = [[2, 3, 5, 0], 11, 13], s=10
	println(sum(primes[2:5]))           // [5, 0, 11] l=3, c=4 -> primes[a] = [2, 3, [5, 0, 11, 0]], s=5
	fmt.Println(primes)                 // primes=[2, 3, 5, 0, 11, 0]

	a = make([]int, 7)
	a = append(a, 13)
	fmt.Printf("%d %d\n", len(a), cap(a))

	// append all slice in other
	b = []int{1, 2, 3, 5, 8}
	a = append(a, b...)

	c := []int{1, 2, 3, 4, 5}
	c = append(c[2:], c[:2]...)
	fmt.Println(c)

	// array into slice
	m := [3]int{1, 2, 3}
	s := m[:]
	fmt.Println(s)
}

func sum(a []int) int {
	a = append(a, 0)
	s := 0
	for a[0] > 0 { // >0 -> s = 5
		s += a[0]
		a = a[1:]
	}
	return s
}
